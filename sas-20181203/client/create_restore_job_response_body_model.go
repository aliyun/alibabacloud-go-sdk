// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRestoreJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateRestoreJobResponseBody
	GetRequestId() *string
}

type CreateRestoreJobResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// BD4B24CE-E5C4-5727-B731-BE85F1D4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRestoreJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRestoreJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRestoreJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRestoreJobResponseBody) SetRequestId(v string) *CreateRestoreJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRestoreJobResponseBody) Validate() error {
	return dara.Validate(s)
}
