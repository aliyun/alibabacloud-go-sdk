// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDownLoadFileUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetDownLoadFileUrlResponseBody
	GetRequestId() *string
	SetCode(v string) *GetDownLoadFileUrlResponseBody
	GetCode() *string
	SetData(v bool) *GetDownLoadFileUrlResponseBody
	GetData() *bool
}

type GetDownLoadFileUrlResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *bool   `json:"data,omitempty" xml:"data,omitempty"`
}

func (s GetDownLoadFileUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDownLoadFileUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetDownLoadFileUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDownLoadFileUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDownLoadFileUrlResponseBody) GetData() *bool {
	return s.Data
}

func (s *GetDownLoadFileUrlResponseBody) SetRequestId(v string) *GetDownLoadFileUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDownLoadFileUrlResponseBody) SetCode(v string) *GetDownLoadFileUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetDownLoadFileUrlResponseBody) SetData(v bool) *GetDownLoadFileUrlResponseBody {
	s.Data = &v
	return s
}

func (s *GetDownLoadFileUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
