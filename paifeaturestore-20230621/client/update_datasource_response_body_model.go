// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDatasourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDatasourceResponseBody
	GetRequestId() *string
}

type UpdateDatasourceResponseBody struct {
	// example:
	//
	// C33E160C-BFCA-5719-B958-942850E949F6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDatasourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasourceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDatasourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDatasourceResponseBody) SetRequestId(v string) *UpdateDatasourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDatasourceResponseBody) Validate() error {
	return dara.Validate(s)
}
