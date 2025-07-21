// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDefaultKmsInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultKmsInstanceId(v string) *GetDefaultKmsInstanceResponseBody
	GetDefaultKmsInstanceId() *string
	SetRequestId(v string) *GetDefaultKmsInstanceResponseBody
	GetRequestId() *string
}

type GetDefaultKmsInstanceResponseBody struct {
	DefaultKmsInstanceId *string `json:"DefaultKmsInstanceId,omitempty" xml:"DefaultKmsInstanceId,omitempty"`
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDefaultKmsInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDefaultKmsInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetDefaultKmsInstanceResponseBody) GetDefaultKmsInstanceId() *string {
	return s.DefaultKmsInstanceId
}

func (s *GetDefaultKmsInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDefaultKmsInstanceResponseBody) SetDefaultKmsInstanceId(v string) *GetDefaultKmsInstanceResponseBody {
	s.DefaultKmsInstanceId = &v
	return s
}

func (s *GetDefaultKmsInstanceResponseBody) SetRequestId(v string) *GetDefaultKmsInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDefaultKmsInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
