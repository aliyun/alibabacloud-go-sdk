// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseDomainSlsConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CloseDomainSlsConfigResponseBody
	GetRequestId() *string
}

type CloseDomainSlsConfigResponseBody struct {
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloseDomainSlsConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloseDomainSlsConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CloseDomainSlsConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloseDomainSlsConfigResponseBody) SetRequestId(v string) *CloseDomainSlsConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseDomainSlsConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
