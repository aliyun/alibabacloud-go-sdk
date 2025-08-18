// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateVersionManagementResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ActivateVersionManagementResponseBody
	GetRequestId() *string
}

type ActivateVersionManagementResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// C370DAF1-C838-4288-A1A0-9A87633D2***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ActivateVersionManagementResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ActivateVersionManagementResponseBody) GoString() string {
	return s.String()
}

func (s *ActivateVersionManagementResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ActivateVersionManagementResponseBody) SetRequestId(v string) *ActivateVersionManagementResponseBody {
	s.RequestId = &v
	return s
}

func (s *ActivateVersionManagementResponseBody) Validate() error {
	return dara.Validate(s)
}
