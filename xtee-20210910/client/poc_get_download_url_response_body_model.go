// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPocGetDownloadUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PocGetDownloadUrlResponseBody
	GetCode() *string
	SetData(v string) *PocGetDownloadUrlResponseBody
	GetData() *string
	SetHttpStatusCode(v string) *PocGetDownloadUrlResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *PocGetDownloadUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *PocGetDownloadUrlResponseBody
	GetRequestId() *string
	SetSuccess(v string) *PocGetDownloadUrlResponseBody
	GetSuccess() *string
}

type PocGetDownloadUrlResponseBody struct {
	// Status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data
	//
	// example:
	//
	// {\\"searchResult\\": [], \\"searchTotalNum\\": 0, \\"searchTime\\": 0.012349}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Response message
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// ID of the request
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. true: Call succeeded. false: Call failed.
	//
	// example:
	//
	// True
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PocGetDownloadUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PocGetDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *PocGetDownloadUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *PocGetDownloadUrlResponseBody) GetData() *string {
	return s.Data
}

func (s *PocGetDownloadUrlResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *PocGetDownloadUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PocGetDownloadUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PocGetDownloadUrlResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *PocGetDownloadUrlResponseBody) SetCode(v string) *PocGetDownloadUrlResponseBody {
	s.Code = &v
	return s
}

func (s *PocGetDownloadUrlResponseBody) SetData(v string) *PocGetDownloadUrlResponseBody {
	s.Data = &v
	return s
}

func (s *PocGetDownloadUrlResponseBody) SetHttpStatusCode(v string) *PocGetDownloadUrlResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PocGetDownloadUrlResponseBody) SetMessage(v string) *PocGetDownloadUrlResponseBody {
	s.Message = &v
	return s
}

func (s *PocGetDownloadUrlResponseBody) SetRequestId(v string) *PocGetDownloadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *PocGetDownloadUrlResponseBody) SetSuccess(v string) *PocGetDownloadUrlResponseBody {
	s.Success = &v
	return s
}

func (s *PocGetDownloadUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
