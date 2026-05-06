// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadSubTaskResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DownloadSubTaskResultResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *DownloadSubTaskResultResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *DownloadSubTaskResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *DownloadSubTaskResultResponseBody
	GetRequestId() *string
	SetResultObject(v *DownloadSubTaskResultResponseBodyResultObject) *DownloadSubTaskResultResponseBody
	GetResultObject() *DownloadSubTaskResultResponseBodyResultObject
}

type DownloadSubTaskResultResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// E01E1B4A-6747-5329-9046-B6D6B2D91349
	RequestId    *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *DownloadSubTaskResultResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s DownloadSubTaskResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DownloadSubTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadSubTaskResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *DownloadSubTaskResultResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DownloadSubTaskResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DownloadSubTaskResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DownloadSubTaskResultResponseBody) GetResultObject() *DownloadSubTaskResultResponseBodyResultObject {
	return s.ResultObject
}

func (s *DownloadSubTaskResultResponseBody) SetCode(v string) *DownloadSubTaskResultResponseBody {
	s.Code = &v
	return s
}

func (s *DownloadSubTaskResultResponseBody) SetHttpStatusCode(v string) *DownloadSubTaskResultResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DownloadSubTaskResultResponseBody) SetMessage(v string) *DownloadSubTaskResultResponseBody {
	s.Message = &v
	return s
}

func (s *DownloadSubTaskResultResponseBody) SetRequestId(v string) *DownloadSubTaskResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DownloadSubTaskResultResponseBody) SetResultObject(v *DownloadSubTaskResultResponseBodyResultObject) *DownloadSubTaskResultResponseBody {
	s.ResultObject = v
	return s
}

func (s *DownloadSubTaskResultResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DownloadSubTaskResultResponseBodyResultObject struct {
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// https://mass.alipay.com/enhance/afts/file/n5XnQounknwAAAAAZfAAAAgAhvocAAFr?t=2hrPX0at3hhaRjlScory9JzLGiLchaonac5suH-Z1BgDAAAAZAABHPpobI2j
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DownloadSubTaskResultResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DownloadSubTaskResultResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DownloadSubTaskResultResponseBodyResultObject) GetStatus() *string {
	return s.Status
}

func (s *DownloadSubTaskResultResponseBodyResultObject) GetUrl() *string {
	return s.Url
}

func (s *DownloadSubTaskResultResponseBodyResultObject) SetStatus(v string) *DownloadSubTaskResultResponseBodyResultObject {
	s.Status = &v
	return s
}

func (s *DownloadSubTaskResultResponseBodyResultObject) SetUrl(v string) *DownloadSubTaskResultResponseBodyResultObject {
	s.Url = &v
	return s
}

func (s *DownloadSubTaskResultResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
