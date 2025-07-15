// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBackendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackendId(v string) *CreateBackendResponseBody
	GetBackendId() *string
	SetRequestId(v string) *CreateBackendResponseBody
	GetRequestId() *string
}

type CreateBackendResponseBody struct {
	// The ID of the backend service.
	//
	// example:
	//
	// 0d105f80a8f340408bd34954d4e4ff22
	BackendId *string `json:"BackendId,omitempty" xml:"BackendId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 66D84355-164D-53ED-81FF-03DCF181DE24
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBackendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBackendResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBackendResponseBody) GetBackendId() *string {
	return s.BackendId
}

func (s *CreateBackendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBackendResponseBody) SetBackendId(v string) *CreateBackendResponseBody {
	s.BackendId = &v
	return s
}

func (s *CreateBackendResponseBody) SetRequestId(v string) *CreateBackendResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBackendResponseBody) Validate() error {
	return dara.Validate(s)
}
