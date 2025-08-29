// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLaboratoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateLaboratoryResponseBody
	GetRequestId() *string
}

type UpdateLaboratoryResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// A04CB8C0-E74A-5E83-BC61-64D153574EC7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLaboratoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLaboratoryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLaboratoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLaboratoryResponseBody) SetRequestId(v string) *UpdateLaboratoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLaboratoryResponseBody) Validate() error {
	return dara.Validate(s)
}
