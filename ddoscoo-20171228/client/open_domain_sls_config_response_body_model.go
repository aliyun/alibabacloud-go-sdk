// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenDomainSlsConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OpenDomainSlsConfigResponseBody
	GetRequestId() *string
}

type OpenDomainSlsConfigResponseBody struct {
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenDomainSlsConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenDomainSlsConfigResponseBody) GoString() string {
	return s.String()
}

func (s *OpenDomainSlsConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenDomainSlsConfigResponseBody) SetRequestId(v string) *OpenDomainSlsConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenDomainSlsConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
