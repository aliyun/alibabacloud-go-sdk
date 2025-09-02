// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDIJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDIJobResponseBody
	GetRequestId() *string
}

type UpdateDIJobResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// AAC30B35-820D-5F3E-A42C-E96BB6379325
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDIJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDIJobResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDIJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDIJobResponseBody) SetRequestId(v string) *UpdateDIJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDIJobResponseBody) Validate() error {
	return dara.Validate(s)
}
