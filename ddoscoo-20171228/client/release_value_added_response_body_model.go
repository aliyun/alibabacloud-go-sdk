// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseValueAddedResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReleaseValueAddedResponseBody
	GetRequestId() *string
}

type ReleaseValueAddedResponseBody struct {
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseValueAddedResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseValueAddedResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseValueAddedResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseValueAddedResponseBody) SetRequestId(v string) *ReleaseValueAddedResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseValueAddedResponseBody) Validate() error {
	return dara.Validate(s)
}
