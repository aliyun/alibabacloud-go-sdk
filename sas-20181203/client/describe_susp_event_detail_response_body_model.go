// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSuspEventDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCanBeDealOnLine(v bool) *DescribeSuspEventDetailResponseBody
	GetCanBeDealOnLine() *bool
	SetDataSource(v string) *DescribeSuspEventDetailResponseBody
	GetDataSource() *string
	SetDetails(v []*DescribeSuspEventDetailResponseBodyDetails) *DescribeSuspEventDetailResponseBody
	GetDetails() []*DescribeSuspEventDetailResponseBodyDetails
	SetEventDesc(v string) *DescribeSuspEventDetailResponseBody
	GetEventDesc() *string
	SetEventName(v string) *DescribeSuspEventDetailResponseBody
	GetEventName() *string
	SetEventStatus(v string) *DescribeSuspEventDetailResponseBody
	GetEventStatus() *string
	SetEventTypeDesc(v string) *DescribeSuspEventDetailResponseBody
	GetEventTypeDesc() *string
	SetId(v int32) *DescribeSuspEventDetailResponseBody
	GetId() *int32
	SetInstanceName(v string) *DescribeSuspEventDetailResponseBody
	GetInstanceName() *string
	SetInternetIp(v string) *DescribeSuspEventDetailResponseBody
	GetInternetIp() *string
	SetIntranetIp(v string) *DescribeSuspEventDetailResponseBody
	GetIntranetIp() *string
	SetLastTime(v string) *DescribeSuspEventDetailResponseBody
	GetLastTime() *string
	SetLevel(v string) *DescribeSuspEventDetailResponseBody
	GetLevel() *string
	SetOperateErrorCode(v string) *DescribeSuspEventDetailResponseBody
	GetOperateErrorCode() *string
	SetOperateMsg(v string) *DescribeSuspEventDetailResponseBody
	GetOperateMsg() *string
	SetRequestId(v string) *DescribeSuspEventDetailResponseBody
	GetRequestId() *string
	SetSaleVersion(v string) *DescribeSuspEventDetailResponseBody
	GetSaleVersion() *string
	SetUuid(v string) *DescribeSuspEventDetailResponseBody
	GetUuid() *string
}

type DescribeSuspEventDetailResponseBody struct {
	// Indicates whether the online processing of exceptions is supported, such as blocking an exception, adding an exception to the whitelist, and ignoring an exception. Valid values:
	//
	// 	- **true**: The online processing of exceptions is supported.
	//
	// 	- **false**: The online processing of exceptions is not supported.
	//
	// example:
	//
	// true
	CanBeDealOnLine *bool `json:"CanBeDealOnLine,omitempty" xml:"CanBeDealOnLine,omitempty"`
	// The data source of the exception.
	//
	// example:
	//
	// aegis_suspicious_****
	DataSource *string `json:"DataSource,omitempty" xml:"DataSource,omitempty"`
	// An array that consists of the details of the exception.
	Details []*DescribeSuspEventDetailResponseBodyDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	// The description of the exception.
	//
	// example:
	//
	// The detection model found a suspicious Webshell file on your server, which may be a backdoor file implanted to maintain permissions after the attacker successfully invaded the website.
	EventDesc *string `json:"EventDesc,omitempty" xml:"EventDesc,omitempty"`
	// The name of the exception.
	//
	// example:
	//
	// WEBSHELL
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The status of the exception. Valid values:
	//
	// 	- **1**: pending handling
	//
	// 	- **2**: ignored
	//
	// 	- **4**: confirmed
	//
	// 	- **8**: marked as a false positive
	//
	// 	- **16**: handling
	//
	// 	- **32**: handled
	//
	// 	- **64**: expired
	//
	// example:
	//
	// 1
	EventStatus *string `json:"EventStatus,omitempty" xml:"EventStatus,omitempty"`
	// The type of the exception.
	//
	// example:
	//
	// Malicious Software-Variable Trojan
	EventTypeDesc *string `json:"EventTypeDesc,omitempty" xml:"EventTypeDesc,omitempty"`
	// The ID of the exception.
	//
	// example:
	//
	// 11416624
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the server on which the exception was detected.
	//
	// example:
	//
	// ca_cpm_****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the server on which the exception was detected.
	//
	// example:
	//
	// 101.132.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the server on which the exception was detected.
	//
	// example:
	//
	// 172.26.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The time when the exception was last detected.
	//
	// example:
	//
	// 2018-10-30 11:43:46
	LastTime *string `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
	// The risk level of the exception. Valid values:
	//
	// 	- **serious**
	//
	// 	- **suspicious**
	//
	// 	- **remind**
	//
	// example:
	//
	// serious
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The code that indicates the handling result of the exception.
	//
	// example:
	//
	// quara.Succes
	OperateErrorCode *string `json:"OperateErrorCode,omitempty" xml:"OperateErrorCode,omitempty"`
	// The message that indicates the handling result of the exception.
	//
	// example:
	//
	// success
	OperateMsg *string `json:"OperateMsg,omitempty" xml:"OperateMsg,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0B48AB3C-84FC-424D-A01D-B9270EF46038
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The edition of Security Center in which the exception can be detected. Valid values:
	//
	// 	- **0**: Basic edition
	//
	// 	- **1**: Advanced edition
	//
	// 	- **2**: Enterprise edition
	//
	// example:
	//
	// 1
	SaleVersion *string `json:"SaleVersion,omitempty" xml:"SaleVersion,omitempty"`
	// The UUID of the server on which the exception was detected.
	//
	// example:
	//
	// bffb12c3-590a-4db2-b538-****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeSuspEventDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSuspEventDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventDetailResponseBody) GetCanBeDealOnLine() *bool {
	return s.CanBeDealOnLine
}

func (s *DescribeSuspEventDetailResponseBody) GetDataSource() *string {
	return s.DataSource
}

func (s *DescribeSuspEventDetailResponseBody) GetDetails() []*DescribeSuspEventDetailResponseBodyDetails {
	return s.Details
}

func (s *DescribeSuspEventDetailResponseBody) GetEventDesc() *string {
	return s.EventDesc
}

func (s *DescribeSuspEventDetailResponseBody) GetEventName() *string {
	return s.EventName
}

func (s *DescribeSuspEventDetailResponseBody) GetEventStatus() *string {
	return s.EventStatus
}

func (s *DescribeSuspEventDetailResponseBody) GetEventTypeDesc() *string {
	return s.EventTypeDesc
}

func (s *DescribeSuspEventDetailResponseBody) GetId() *int32 {
	return s.Id
}

func (s *DescribeSuspEventDetailResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeSuspEventDetailResponseBody) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeSuspEventDetailResponseBody) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribeSuspEventDetailResponseBody) GetLastTime() *string {
	return s.LastTime
}

func (s *DescribeSuspEventDetailResponseBody) GetLevel() *string {
	return s.Level
}

func (s *DescribeSuspEventDetailResponseBody) GetOperateErrorCode() *string {
	return s.OperateErrorCode
}

func (s *DescribeSuspEventDetailResponseBody) GetOperateMsg() *string {
	return s.OperateMsg
}

func (s *DescribeSuspEventDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSuspEventDetailResponseBody) GetSaleVersion() *string {
	return s.SaleVersion
}

func (s *DescribeSuspEventDetailResponseBody) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeSuspEventDetailResponseBody) SetCanBeDealOnLine(v bool) *DescribeSuspEventDetailResponseBody {
	s.CanBeDealOnLine = &v
	return s
}

func (s *DescribeSuspEventDetailResponseBody) SetDataSource(v string) *DescribeSuspEventDetailResponseBody {
	s.DataSource = &v
	return s
}

func (s *DescribeSuspEventDetailResponseBody) SetDetails(v []*DescribeSuspEventDetailResponseBodyDetails) *DescribeSuspEventDetailResponseBody {
	s.Details = v
	return s
}

func (s *DescribeSuspEventDetailResponseBody) SetEventDesc(v string) *DescribeSuspEventDetailResponseBody {
	s.EventDesc = &v
	return s
}

func (s *DescribeSuspEventDetailResponseBody) SetEventName(v string) *DescribeSuspEventDetailResponseBody {
	s.EventName = &v
	return s
}

func (s *DescribeSuspEventDetailResponseBody) SetEventStatus(v string) *DescribeSuspEventDetailResponseBody {
	s.EventStatus = &v
	return s
}

func (s *DescribeSuspEventDetailResponseBody) SetEventTypeDesc(v string) *DescribeSuspEventDetailResponseBody {
	s.EventTypeDesc = &v
	return s
}

func (s *DescribeSuspEventDetailResponseBody) SetId(v int32) *DescribeSuspEventDetailResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeSuspEventDetailResponseBody) SetInstanceName(v string) *DescribeSuspEventDetailResponseBody {
	s.InstanceName = &v
	return s
}

func (s *DescribeSuspEventDetailResponseBody) SetInternetIp(v string) *DescribeSuspEventDetailResponseBody {
	s.InternetIp = &v
	return s
}

func (s *DescribeSuspEventDetailResponseBody) SetIntranetIp(v string) *DescribeSuspEventDetailResponseBody {
	s.IntranetIp = &v
	return s
}

func (s *DescribeSuspEventDetailResponseBody) SetLastTime(v string) *DescribeSuspEventDetailResponseBody {
	s.LastTime = &v
	return s
}

func (s *DescribeSuspEventDetailResponseBody) SetLevel(v string) *DescribeSuspEventDetailResponseBody {
	s.Level = &v
	return s
}

func (s *DescribeSuspEventDetailResponseBody) SetOperateErrorCode(v string) *DescribeSuspEventDetailResponseBody {
	s.OperateErrorCode = &v
	return s
}

func (s *DescribeSuspEventDetailResponseBody) SetOperateMsg(v string) *DescribeSuspEventDetailResponseBody {
	s.OperateMsg = &v
	return s
}

func (s *DescribeSuspEventDetailResponseBody) SetRequestId(v string) *DescribeSuspEventDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSuspEventDetailResponseBody) SetSaleVersion(v string) *DescribeSuspEventDetailResponseBody {
	s.SaleVersion = &v
	return s
}

func (s *DescribeSuspEventDetailResponseBody) SetUuid(v string) *DescribeSuspEventDetailResponseBody {
	s.Uuid = &v
	return s
}

func (s *DescribeSuspEventDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSuspEventDetailResponseBodyDetails struct {
	// The display name of the alert event.
	//
	// example:
	//
	// Trojan Path
	NameDisplay *string `json:"NameDisplay,omitempty" xml:"NameDisplay,omitempty"`
	// The format in which the details of the exception are displayed.
	//
	// Valid values:
	//
	// 	- **text**
	//
	// 	- **html**
	//
	// example:
	//
	// html
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The attribute information about the exception. For example, if the exception is associated with an alert that is triggered by an unusual logon, the information can include the time when the logon is initiated and the location from which the logon is initiated. If the exception is associated with an alert that is triggered by a webshell file, the information can include the path of the trojan file and the type of the trojan.
	//
	// example:
	//
	// getopt
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSuspEventDetailResponseBodyDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeSuspEventDetailResponseBodyDetails) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventDetailResponseBodyDetails) GetNameDisplay() *string {
	return s.NameDisplay
}

func (s *DescribeSuspEventDetailResponseBodyDetails) GetType() *string {
	return s.Type
}

func (s *DescribeSuspEventDetailResponseBodyDetails) GetValue() *string {
	return s.Value
}

func (s *DescribeSuspEventDetailResponseBodyDetails) SetNameDisplay(v string) *DescribeSuspEventDetailResponseBodyDetails {
	s.NameDisplay = &v
	return s
}

func (s *DescribeSuspEventDetailResponseBodyDetails) SetType(v string) *DescribeSuspEventDetailResponseBodyDetails {
	s.Type = &v
	return s
}

func (s *DescribeSuspEventDetailResponseBodyDetails) SetValue(v string) *DescribeSuspEventDetailResponseBodyDetails {
	s.Value = &v
	return s
}

func (s *DescribeSuspEventDetailResponseBodyDetails) Validate() error {
	return dara.Validate(s)
}
