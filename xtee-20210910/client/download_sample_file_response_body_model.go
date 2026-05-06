// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadSampleFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DownloadSampleFileResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *DownloadSampleFileResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *DownloadSampleFileResponseBody
	GetMessage() *string
	SetRequestId(v string) *DownloadSampleFileResponseBody
	GetRequestId() *string
	SetResultObject(v *DownloadSampleFileResponseBodyResultObject) *DownloadSampleFileResponseBody
	GetResultObject() *DownloadSampleFileResponseBodyResultObject
}

type DownloadSampleFileResponseBody struct {
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
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *DownloadSampleFileResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s DownloadSampleFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DownloadSampleFileResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadSampleFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *DownloadSampleFileResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DownloadSampleFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DownloadSampleFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DownloadSampleFileResponseBody) GetResultObject() *DownloadSampleFileResponseBodyResultObject {
	return s.ResultObject
}

func (s *DownloadSampleFileResponseBody) SetCode(v string) *DownloadSampleFileResponseBody {
	s.Code = &v
	return s
}

func (s *DownloadSampleFileResponseBody) SetHttpStatusCode(v string) *DownloadSampleFileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DownloadSampleFileResponseBody) SetMessage(v string) *DownloadSampleFileResponseBody {
	s.Message = &v
	return s
}

func (s *DownloadSampleFileResponseBody) SetRequestId(v string) *DownloadSampleFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DownloadSampleFileResponseBody) SetResultObject(v *DownloadSampleFileResponseBodyResultObject) *DownloadSampleFileResponseBody {
	s.ResultObject = v
	return s
}

func (s *DownloadSampleFileResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DownloadSampleFileResponseBodyResultObject struct {
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// https://pic.zfp.cn/image/2026/02/06/45b5fee18baf4b99b13025987486319c.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DownloadSampleFileResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DownloadSampleFileResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DownloadSampleFileResponseBodyResultObject) GetStatus() *string {
	return s.Status
}

func (s *DownloadSampleFileResponseBodyResultObject) GetUrl() *string {
	return s.Url
}

func (s *DownloadSampleFileResponseBodyResultObject) SetStatus(v string) *DownloadSampleFileResponseBodyResultObject {
	s.Status = &v
	return s
}

func (s *DownloadSampleFileResponseBodyResultObject) SetUrl(v string) *DownloadSampleFileResponseBodyResultObject {
	s.Url = &v
	return s
}

func (s *DownloadSampleFileResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
