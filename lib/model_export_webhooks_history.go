/*
 * Brevo API
 *
 * Brevo provide a RESTFul API that can be used with any languages. With this API, you will be able to :   - Manage your campaigns and get the statistics   - Manage your contacts   - Send transactional Emails and SMS   - and much more...  You can download our wrappers at https://github.com/orgs/brevo  **Possible responses**   | Code | Message |   | :-------------: | ------------- |   | 200  | OK. Successful Request  |   | 201  | OK. Successful Creation |   | 202  | OK. Request accepted |   | 204  | OK. Successful Update/Deletion  |   | 400  | Error. Bad Request  |   | 401  | Error. Authentication Needed  |   | 402  | Error. Not enough credit, plan upgrade needed  |   | 403  | Error. Permission denied  |   | 404  | Error. Object does not exist |   | 405  | Error. Method not allowed  |   | 406  | Error. Not Acceptable  |
 *
 * API version: 3.0.0
 * Contact: contact@brevo.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ExportWebhooksHistory struct {
	// Number of days in the past including today (positive integer). _Not compatible with 'startDate' and 'endDate'_
	Days int32 `json:"days,omitempty"`
	// Mandatory if endDate is used. Starting date of the history (YYYY-MM-DD). Must be lower than equal to endDate
	StartDate string `json:"startDate,omitempty"`
	// Mandatory if startDate is used. Ending date of the report (YYYY-MM-DD). Must be greater than equal to startDate
	EndDate string `json:"endDate,omitempty"`
	// Sorting order of records (asc or desc)
	Sort string `json:"sort,omitempty"`
	// Filter the history for a specific event type
	Event string `json:"event"`
	// Webhook URL to receive CSV file link
	NotifyURL string `json:"notifyURL"`
	// Filter the history for a specific webhook id
	WebhookId int32 `json:"webhookId,omitempty"`
	// Filter the history for a specific email
	Email string `json:"email,omitempty"`
	// Filter the history for a specific message id
	MessageId int32 `json:"messageId,omitempty"`
}