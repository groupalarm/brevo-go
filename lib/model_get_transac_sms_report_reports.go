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

type GetTransacSmsReportReports struct {
	// Date for which statistics are retrieved
	Date string `json:"date,omitempty"`
	// Number of requests for the date
	Requests int64 `json:"requests,omitempty"`
	// Number of delivered SMS for the date
	Delivered int64 `json:"delivered,omitempty"`
	// Number of hardbounces for the date
	HardBounces int64 `json:"hardBounces,omitempty"`
	// Number of softbounces for the date
	SoftBounces int64 `json:"softBounces,omitempty"`
	// Number of blocked contact for the date
	Blocked int64 `json:"blocked,omitempty"`
	// Number of unsubscription for the date
	Unsubscribed int64 `json:"unsubscribed,omitempty"`
	// Number of answered SMS for the date
	Replied int64 `json:"replied,omitempty"`
	// Number of accepted for the date
	Accepted int64 `json:"accepted,omitempty"`
	// Number of rejected for the date
	Rejected int64 `json:"rejected,omitempty"`
}