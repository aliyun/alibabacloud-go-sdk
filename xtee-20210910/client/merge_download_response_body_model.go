// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMergeDownloadResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MergeDownloadResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *MergeDownloadResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *MergeDownloadResponseBody
	GetMessage() *string
	SetRequestId(v string) *MergeDownloadResponseBody
	GetRequestId() *string
	SetResultObject(v *MergeDownloadResponseBodyResultObject) *MergeDownloadResponseBody
	GetResultObject() *MergeDownloadResponseBodyResultObject
}

type MergeDownloadResponseBody struct {
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
	RequestId    *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *MergeDownloadResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s MergeDownloadResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MergeDownloadResponseBody) GoString() string {
	return s.String()
}

func (s *MergeDownloadResponseBody) GetCode() *string {
	return s.Code
}

func (s *MergeDownloadResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *MergeDownloadResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MergeDownloadResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MergeDownloadResponseBody) GetResultObject() *MergeDownloadResponseBodyResultObject {
	return s.ResultObject
}

func (s *MergeDownloadResponseBody) SetCode(v string) *MergeDownloadResponseBody {
	s.Code = &v
	return s
}

func (s *MergeDownloadResponseBody) SetHttpStatusCode(v string) *MergeDownloadResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *MergeDownloadResponseBody) SetMessage(v string) *MergeDownloadResponseBody {
	s.Message = &v
	return s
}

func (s *MergeDownloadResponseBody) SetRequestId(v string) *MergeDownloadResponseBody {
	s.RequestId = &v
	return s
}

func (s *MergeDownloadResponseBody) SetResultObject(v *MergeDownloadResponseBodyResultObject) *MergeDownloadResponseBody {
	s.ResultObject = v
	return s
}

func (s *MergeDownloadResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MergeDownloadResponseBodyResultObject struct {
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// rtmp://push-live.gaotime.com/jiaenguang789/21645251932448?auth_key=1768357031-0-0-6e2a4815fe4e2287a86be1105d42bf53
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s MergeDownloadResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s MergeDownloadResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *MergeDownloadResponseBodyResultObject) GetStatus() *string {
	return s.Status
}

func (s *MergeDownloadResponseBodyResultObject) GetUrl() *string {
	return s.Url
}

func (s *MergeDownloadResponseBodyResultObject) SetStatus(v string) *MergeDownloadResponseBodyResultObject {
	s.Status = &v
	return s
}

func (s *MergeDownloadResponseBodyResultObject) SetUrl(v string) *MergeDownloadResponseBodyResultObject {
	s.Url = &v
	return s
}

func (s *MergeDownloadResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
