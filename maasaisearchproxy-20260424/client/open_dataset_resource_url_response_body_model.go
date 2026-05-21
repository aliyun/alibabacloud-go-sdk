// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenDatasetResourceUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *OpenDatasetResourceUrlResponseBody
	GetCode() *int32
	SetData(v *OpenDatasetResourceUrlResponseBodyData) *OpenDatasetResourceUrlResponseBody
	GetData() *OpenDatasetResourceUrlResponseBodyData
	SetMessage(v string) *OpenDatasetResourceUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *OpenDatasetResourceUrlResponseBody
	GetRequestId() *string
}

type OpenDatasetResourceUrlResponseBody struct {
	// example:
	//
	// 200
	Code *int32                                  `json:"code,omitempty" xml:"code,omitempty"`
	Data *OpenDatasetResourceUrlResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 1a0f40dd17774641794394269ec0e9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s OpenDatasetResourceUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenDatasetResourceUrlResponseBody) GoString() string {
	return s.String()
}

func (s *OpenDatasetResourceUrlResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *OpenDatasetResourceUrlResponseBody) GetData() *OpenDatasetResourceUrlResponseBodyData {
	return s.Data
}

func (s *OpenDatasetResourceUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OpenDatasetResourceUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenDatasetResourceUrlResponseBody) SetCode(v int32) *OpenDatasetResourceUrlResponseBody {
	s.Code = &v
	return s
}

func (s *OpenDatasetResourceUrlResponseBody) SetData(v *OpenDatasetResourceUrlResponseBodyData) *OpenDatasetResourceUrlResponseBody {
	s.Data = v
	return s
}

func (s *OpenDatasetResourceUrlResponseBody) SetMessage(v string) *OpenDatasetResourceUrlResponseBody {
	s.Message = &v
	return s
}

func (s *OpenDatasetResourceUrlResponseBody) SetRequestId(v string) *OpenDatasetResourceUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenDatasetResourceUrlResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OpenDatasetResourceUrlResponseBodyData struct {
	// example:
	//
	// https://maas-ai-search-center-raw.oss-cn-hangzhou.aliyuncs.com/.../sample.mp4...
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s OpenDatasetResourceUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s OpenDatasetResourceUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *OpenDatasetResourceUrlResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *OpenDatasetResourceUrlResponseBodyData) SetUrl(v string) *OpenDatasetResourceUrlResponseBodyData {
	s.Url = &v
	return s
}

func (s *OpenDatasetResourceUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
