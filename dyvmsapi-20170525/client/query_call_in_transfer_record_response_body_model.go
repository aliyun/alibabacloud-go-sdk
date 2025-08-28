// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCallInTransferRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryCallInTransferRecordResponseBody
	GetCode() *string
	SetData(v *QueryCallInTransferRecordResponseBodyData) *QueryCallInTransferRecordResponseBody
	GetData() *QueryCallInTransferRecordResponseBodyData
	SetMessage(v string) *QueryCallInTransferRecordResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryCallInTransferRecordResponseBody
	GetRequestId() *string
}

type QueryCallInTransferRecordResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- For more information about other response codes, see [API error codes](https://help.aliyun.com/document_detail/112502.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *QueryCallInTransferRecordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7BF47617-7851-48F7-A3A1-2021342A78E2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryCallInTransferRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCallInTransferRecordResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCallInTransferRecordResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryCallInTransferRecordResponseBody) GetData() *QueryCallInTransferRecordResponseBodyData {
	return s.Data
}

func (s *QueryCallInTransferRecordResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryCallInTransferRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCallInTransferRecordResponseBody) SetCode(v string) *QueryCallInTransferRecordResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCallInTransferRecordResponseBody) SetData(v *QueryCallInTransferRecordResponseBodyData) *QueryCallInTransferRecordResponseBody {
	s.Data = v
	return s
}

func (s *QueryCallInTransferRecordResponseBody) SetMessage(v string) *QueryCallInTransferRecordResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCallInTransferRecordResponseBody) SetRequestId(v string) *QueryCallInTransferRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCallInTransferRecordResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryCallInTransferRecordResponseBodyData struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 50
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
	// The call transfer records.
	Values []*QueryCallInTransferRecordResponseBodyDataValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s QueryCallInTransferRecordResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryCallInTransferRecordResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryCallInTransferRecordResponseBodyData) GetPageNo() *int64 {
	return s.PageNo
}

func (s *QueryCallInTransferRecordResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QueryCallInTransferRecordResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *QueryCallInTransferRecordResponseBodyData) GetValues() []*QueryCallInTransferRecordResponseBodyDataValues {
	return s.Values
}

func (s *QueryCallInTransferRecordResponseBodyData) SetPageNo(v int64) *QueryCallInTransferRecordResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *QueryCallInTransferRecordResponseBodyData) SetPageSize(v int64) *QueryCallInTransferRecordResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryCallInTransferRecordResponseBodyData) SetTotal(v int64) *QueryCallInTransferRecordResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryCallInTransferRecordResponseBodyData) SetValues(v []*QueryCallInTransferRecordResponseBodyDataValues) *QueryCallInTransferRecordResponseBodyData {
	s.Values = v
	return s
}

func (s *QueryCallInTransferRecordResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type QueryCallInTransferRecordResponseBodyDataValues struct {
	// The called number of the inbound call.
	//
	// example:
	//
	// 150****0000
	CallInCalled *string `json:"CallInCalled,omitempty" xml:"CallInCalled,omitempty"`
	// The calling number of the inbound call.
	//
	// example:
	//
	// 150****0000
	CallInCaller *string `json:"CallInCaller,omitempty" xml:"CallInCaller,omitempty"`
	// The time when the call was initiated.
	//
	// example:
	//
	// 2020-10-03 10:21:21
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The recording URL.
	//
	// example:
	//
	// http://alicom-fc-media.cn-hangzhou.oss.aliyun-inc.com/dayuBizVoiceMrg/4238c230-9e74-41be-90b8-2fbe7684****.wav?Expires=1627538265&OSSAccessKeyId=bypFNbGJVk73****&Signature=****mUqkPqIQ%3D
	RecordUrl *string `json:"RecordUrl,omitempty" xml:"RecordUrl,omitempty"`
	// The phone number to which the call was transferred.
	//
	// example:
	//
	// 151****0000
	TransferCalled *string `json:"TransferCalled,omitempty" xml:"TransferCalled,omitempty"`
	// The calling number that transferred the call.
	//
	// example:
	//
	// 151****0000
	TransferCaller *string `json:"TransferCaller,omitempty" xml:"TransferCaller,omitempty"`
}

func (s QueryCallInTransferRecordResponseBodyDataValues) String() string {
	return dara.Prettify(s)
}

func (s QueryCallInTransferRecordResponseBodyDataValues) GoString() string {
	return s.String()
}

func (s *QueryCallInTransferRecordResponseBodyDataValues) GetCallInCalled() *string {
	return s.CallInCalled
}

func (s *QueryCallInTransferRecordResponseBodyDataValues) GetCallInCaller() *string {
	return s.CallInCaller
}

func (s *QueryCallInTransferRecordResponseBodyDataValues) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *QueryCallInTransferRecordResponseBodyDataValues) GetRecordUrl() *string {
	return s.RecordUrl
}

func (s *QueryCallInTransferRecordResponseBodyDataValues) GetTransferCalled() *string {
	return s.TransferCalled
}

func (s *QueryCallInTransferRecordResponseBodyDataValues) GetTransferCaller() *string {
	return s.TransferCaller
}

func (s *QueryCallInTransferRecordResponseBodyDataValues) SetCallInCalled(v string) *QueryCallInTransferRecordResponseBodyDataValues {
	s.CallInCalled = &v
	return s
}

func (s *QueryCallInTransferRecordResponseBodyDataValues) SetCallInCaller(v string) *QueryCallInTransferRecordResponseBodyDataValues {
	s.CallInCaller = &v
	return s
}

func (s *QueryCallInTransferRecordResponseBodyDataValues) SetGmtCreate(v int64) *QueryCallInTransferRecordResponseBodyDataValues {
	s.GmtCreate = &v
	return s
}

func (s *QueryCallInTransferRecordResponseBodyDataValues) SetRecordUrl(v string) *QueryCallInTransferRecordResponseBodyDataValues {
	s.RecordUrl = &v
	return s
}

func (s *QueryCallInTransferRecordResponseBodyDataValues) SetTransferCalled(v string) *QueryCallInTransferRecordResponseBodyDataValues {
	s.TransferCalled = &v
	return s
}

func (s *QueryCallInTransferRecordResponseBodyDataValues) SetTransferCaller(v string) *QueryCallInTransferRecordResponseBodyDataValues {
	s.TransferCaller = &v
	return s
}

func (s *QueryCallInTransferRecordResponseBodyDataValues) Validate() error {
	return dara.Validate(s)
}
