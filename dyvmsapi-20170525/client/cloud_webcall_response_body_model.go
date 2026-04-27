// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudWebcallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudWebcallResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudWebcallResponseBody
	GetCode() *string
	SetData(v *CloudWebcallResponseBodyData) *CloudWebcallResponseBody
	GetData() *CloudWebcallResponseBodyData
	SetMessage(v string) *CloudWebcallResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudWebcallResponseBody
	GetRequestId() *string
}

type CloudWebcallResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudWebcallResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9FF70B74-1B3C-44C1-ACDF-8DF962988F0E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudWebcallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudWebcallResponseBody) GoString() string {
	return s.String()
}

func (s *CloudWebcallResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudWebcallResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudWebcallResponseBody) GetData() *CloudWebcallResponseBodyData {
	return s.Data
}

func (s *CloudWebcallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudWebcallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudWebcallResponseBody) SetAccessDeniedDetail(v string) *CloudWebcallResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudWebcallResponseBody) SetCode(v string) *CloudWebcallResponseBody {
	s.Code = &v
	return s
}

func (s *CloudWebcallResponseBody) SetData(v *CloudWebcallResponseBodyData) *CloudWebcallResponseBody {
	s.Data = v
	return s
}

func (s *CloudWebcallResponseBody) SetMessage(v string) *CloudWebcallResponseBody {
	s.Message = &v
	return s
}

func (s *CloudWebcallResponseBody) SetRequestId(v string) *CloudWebcallResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudWebcallResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudWebcallResponseBodyData struct {
	// 结果
	//
	// example:
	//
	// 0
	Result *int64 `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CloudWebcallResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudWebcallResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudWebcallResponseBodyData) GetResult() *int64 {
	return s.Result
}

func (s *CloudWebcallResponseBodyData) SetResult(v int64) *CloudWebcallResponseBodyData {
	s.Result = &v
	return s
}

func (s *CloudWebcallResponseBodyData) Validate() error {
	return dara.Validate(s)
}
