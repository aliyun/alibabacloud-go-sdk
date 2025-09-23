// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDBClusterOrcaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DisableDBClusterOrcaResponseBody
	GetDBClusterId() *string
	SetRequestId(v string) *DisableDBClusterOrcaResponseBody
	GetRequestId() *string
}

type DisableDBClusterOrcaResponseBody struct {
	// example:
	//
	// pc-***************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// D0CEC6AC-7760-409A-A0D5-E6CD86******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableDBClusterOrcaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableDBClusterOrcaResponseBody) GoString() string {
	return s.String()
}

func (s *DisableDBClusterOrcaResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DisableDBClusterOrcaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableDBClusterOrcaResponseBody) SetDBClusterId(v string) *DisableDBClusterOrcaResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DisableDBClusterOrcaResponseBody) SetRequestId(v string) *DisableDBClusterOrcaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableDBClusterOrcaResponseBody) Validate() error {
	return dara.Validate(s)
}
