// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSampleFileDownloadResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SampleFileDownloadResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *SampleFileDownloadResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *SampleFileDownloadResponseBody
	GetMessage() *string
	SetRequestId(v string) *SampleFileDownloadResponseBody
	GetRequestId() *string
	SetResultObject(v string) *SampleFileDownloadResponseBody
	GetResultObject() *string
}

type SampleFileDownloadResponseBody struct {
	// Status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Return message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Request result.
	//
	// example:
	//
	// true
	ResultObject *string `json:"ResultObject,omitempty" xml:"ResultObject,omitempty"`
}

func (s SampleFileDownloadResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SampleFileDownloadResponseBody) GoString() string {
	return s.String()
}

func (s *SampleFileDownloadResponseBody) GetCode() *string {
	return s.Code
}

func (s *SampleFileDownloadResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *SampleFileDownloadResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SampleFileDownloadResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SampleFileDownloadResponseBody) GetResultObject() *string {
	return s.ResultObject
}

func (s *SampleFileDownloadResponseBody) SetCode(v string) *SampleFileDownloadResponseBody {
	s.Code = &v
	return s
}

func (s *SampleFileDownloadResponseBody) SetHttpStatusCode(v string) *SampleFileDownloadResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SampleFileDownloadResponseBody) SetMessage(v string) *SampleFileDownloadResponseBody {
	s.Message = &v
	return s
}

func (s *SampleFileDownloadResponseBody) SetRequestId(v string) *SampleFileDownloadResponseBody {
	s.RequestId = &v
	return s
}

func (s *SampleFileDownloadResponseBody) SetResultObject(v string) *SampleFileDownloadResponseBody {
	s.ResultObject = &v
	return s
}

func (s *SampleFileDownloadResponseBody) Validate() error {
	return dara.Validate(s)
}
