// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatasetResourceUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetDatasetResourceUrlResponseBody
	GetCode() *int32
	SetData(v *GetDatasetResourceUrlResponseBodyData) *GetDatasetResourceUrlResponseBody
	GetData() *GetDatasetResourceUrlResponseBodyData
	SetMessage(v string) *GetDatasetResourceUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDatasetResourceUrlResponseBody
	GetRequestId() *string
}

type GetDatasetResourceUrlResponseBody struct {
	// The business status code. A value of `200` indicates a successful request. Other values indicate an exception. For more information, see error codes.
	//
	// example:
	//
	// 200
	Code *int32 `json:"code,omitempty" xml:"code,omitempty"`
	// The business data body.
	Data *GetDatasetResourceUrlResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The status description. The value is "success" for successful requests and a specific error message for failed requests.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The unique request ID, used for troubleshooting.
	//
	// example:
	//
	// 1a0f40dd17774641794394269ec0e9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetDatasetResourceUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetResourceUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetDatasetResourceUrlResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetDatasetResourceUrlResponseBody) GetData() *GetDatasetResourceUrlResponseBodyData {
	return s.Data
}

func (s *GetDatasetResourceUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDatasetResourceUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDatasetResourceUrlResponseBody) SetCode(v int32) *GetDatasetResourceUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetDatasetResourceUrlResponseBody) SetData(v *GetDatasetResourceUrlResponseBodyData) *GetDatasetResourceUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetDatasetResourceUrlResponseBody) SetMessage(v string) *GetDatasetResourceUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetDatasetResourceUrlResponseBody) SetRequestId(v string) *GetDatasetResourceUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDatasetResourceUrlResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDatasetResourceUrlResponseBodyData struct {
	// The temporary OSS access URL with a signature and expiration time (valid for 24 hours). The URL can be used directly for frontend display or download.
	//
	// example:
	//
	// https://maas-ai-search-center-raw.oss-cn-hangzhou.aliyuncs.com/.../sample.mp4...
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetDatasetResourceUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetResourceUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDatasetResourceUrlResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *GetDatasetResourceUrlResponseBodyData) SetUrl(v string) *GetDatasetResourceUrlResponseBodyData {
	s.Url = &v
	return s
}

func (s *GetDatasetResourceUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
