// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBackendModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackendModelId(v string) *CreateBackendModelResponseBody
	GetBackendModelId() *string
	SetRequestId(v string) *CreateBackendModelResponseBody
	GetRequestId() *string
}

type CreateBackendModelResponseBody struct {
	// example:
	//
	// 4be6b110b7aa40b0bf0c83cc00b3bd86
	BackendModelId *string `json:"BackendModelId,omitempty" xml:"BackendModelId,omitempty"`
	// example:
	//
	// 64411ECF-FAF7-5E3C-BA7B-E4A1F15A28CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBackendModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBackendModelResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBackendModelResponseBody) GetBackendModelId() *string {
	return s.BackendModelId
}

func (s *CreateBackendModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBackendModelResponseBody) SetBackendModelId(v string) *CreateBackendModelResponseBody {
	s.BackendModelId = &v
	return s
}

func (s *CreateBackendModelResponseBody) SetRequestId(v string) *CreateBackendModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBackendModelResponseBody) Validate() error {
	return dara.Validate(s)
}
