// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudDeleteExtenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudDeleteExtenResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudDeleteExtenResponseBody
	GetCode() *string
	SetData(v *CloudDeleteExtenResponseBodyData) *CloudDeleteExtenResponseBody
	GetData() *CloudDeleteExtenResponseBodyData
	SetMessage(v string) *CloudDeleteExtenResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudDeleteExtenResponseBody
	GetRequestId() *string
}

type CloudDeleteExtenResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudDeleteExtenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 7BF47617-7851-48F7-A3A1-2021342A78E2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudDeleteExtenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudDeleteExtenResponseBody) GoString() string {
	return s.String()
}

func (s *CloudDeleteExtenResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudDeleteExtenResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudDeleteExtenResponseBody) GetData() *CloudDeleteExtenResponseBodyData {
	return s.Data
}

func (s *CloudDeleteExtenResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudDeleteExtenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudDeleteExtenResponseBody) SetAccessDeniedDetail(v string) *CloudDeleteExtenResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudDeleteExtenResponseBody) SetCode(v string) *CloudDeleteExtenResponseBody {
	s.Code = &v
	return s
}

func (s *CloudDeleteExtenResponseBody) SetData(v *CloudDeleteExtenResponseBodyData) *CloudDeleteExtenResponseBody {
	s.Data = v
	return s
}

func (s *CloudDeleteExtenResponseBody) SetMessage(v string) *CloudDeleteExtenResponseBody {
	s.Message = &v
	return s
}

func (s *CloudDeleteExtenResponseBody) SetRequestId(v string) *CloudDeleteExtenResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudDeleteExtenResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudDeleteExtenResponseBodyData struct {
	// example:
	//
	// 0
	Result *int64 `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CloudDeleteExtenResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudDeleteExtenResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudDeleteExtenResponseBodyData) GetResult() *int64 {
	return s.Result
}

func (s *CloudDeleteExtenResponseBodyData) SetResult(v int64) *CloudDeleteExtenResponseBodyData {
	s.Result = &v
	return s
}

func (s *CloudDeleteExtenResponseBodyData) Validate() error {
	return dara.Validate(s)
}
