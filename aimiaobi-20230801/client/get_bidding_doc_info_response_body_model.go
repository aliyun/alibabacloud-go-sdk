// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBiddingDocInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetBiddingDocInfoResponseBody
	GetCode() *string
	SetData(v *GetBiddingDocInfoResponseBodyData) *GetBiddingDocInfoResponseBody
	GetData() *GetBiddingDocInfoResponseBodyData
	SetHttpStatusCode(v int32) *GetBiddingDocInfoResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetBiddingDocInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetBiddingDocInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetBiddingDocInfoResponseBody
	GetSuccess() *bool
}

type GetBiddingDocInfoResponseBody struct {
	// example:
	//
	// successful
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetBiddingDocInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F2F366D6-E9FE-1006-BB70-2C650896AAB5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetBiddingDocInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBiddingDocInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetBiddingDocInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetBiddingDocInfoResponseBody) GetData() *GetBiddingDocInfoResponseBodyData {
	return s.Data
}

func (s *GetBiddingDocInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetBiddingDocInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetBiddingDocInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBiddingDocInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetBiddingDocInfoResponseBody) SetCode(v string) *GetBiddingDocInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetBiddingDocInfoResponseBody) SetData(v *GetBiddingDocInfoResponseBodyData) *GetBiddingDocInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetBiddingDocInfoResponseBody) SetHttpStatusCode(v int32) *GetBiddingDocInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetBiddingDocInfoResponseBody) SetMessage(v string) *GetBiddingDocInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetBiddingDocInfoResponseBody) SetRequestId(v string) *GetBiddingDocInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBiddingDocInfoResponseBody) SetSuccess(v bool) *GetBiddingDocInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetBiddingDocInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetBiddingDocInfoResponseBodyData struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// markdown
	//
	// html
	ContentFormat *string `json:"ContentFormat,omitempty" xml:"ContentFormat,omitempty"`
	// example:
	//
	// outline
	//
	// bidding
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// example:
	//
	// 0-waiting、1-running、2-success、3-pause、4-fail
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// analysis
	//
	// writing
	Step *string `json:"Step,omitempty" xml:"Step,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// http://xxx
	TenderDocUrl *string `json:"TenderDocUrl,omitempty" xml:"TenderDocUrl,omitempty"`
	// example:
	//
	// pdf
	//
	// docx
	TenderFileType *string `json:"TenderFileType,omitempty" xml:"TenderFileType,omitempty"`
}

func (s GetBiddingDocInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetBiddingDocInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetBiddingDocInfoResponseBodyData) GetContent() *string {
	return s.Content
}

func (s *GetBiddingDocInfoResponseBodyData) GetContentFormat() *string {
	return s.ContentFormat
}

func (s *GetBiddingDocInfoResponseBodyData) GetContentType() *string {
	return s.ContentType
}

func (s *GetBiddingDocInfoResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *GetBiddingDocInfoResponseBodyData) GetStep() *string {
	return s.Step
}

func (s *GetBiddingDocInfoResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetBiddingDocInfoResponseBodyData) GetTenderDocUrl() *string {
	return s.TenderDocUrl
}

func (s *GetBiddingDocInfoResponseBodyData) GetTenderFileType() *string {
	return s.TenderFileType
}

func (s *GetBiddingDocInfoResponseBodyData) SetContent(v string) *GetBiddingDocInfoResponseBodyData {
	s.Content = &v
	return s
}

func (s *GetBiddingDocInfoResponseBodyData) SetContentFormat(v string) *GetBiddingDocInfoResponseBodyData {
	s.ContentFormat = &v
	return s
}

func (s *GetBiddingDocInfoResponseBodyData) SetContentType(v string) *GetBiddingDocInfoResponseBodyData {
	s.ContentType = &v
	return s
}

func (s *GetBiddingDocInfoResponseBodyData) SetStatus(v int32) *GetBiddingDocInfoResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetBiddingDocInfoResponseBodyData) SetStep(v string) *GetBiddingDocInfoResponseBodyData {
	s.Step = &v
	return s
}

func (s *GetBiddingDocInfoResponseBodyData) SetTaskId(v string) *GetBiddingDocInfoResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetBiddingDocInfoResponseBodyData) SetTenderDocUrl(v string) *GetBiddingDocInfoResponseBodyData {
	s.TenderDocUrl = &v
	return s
}

func (s *GetBiddingDocInfoResponseBodyData) SetTenderFileType(v string) *GetBiddingDocInfoResponseBodyData {
	s.TenderFileType = &v
	return s
}

func (s *GetBiddingDocInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
