// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnterpriseAccelerateTargetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEapId(v string) *ListEnterpriseAccelerateTargetsResponseBody
	GetEapId() *string
	SetRequestId(v string) *ListEnterpriseAccelerateTargetsResponseBody
	GetRequestId() *string
	SetTargets(v []*string) *ListEnterpriseAccelerateTargetsResponseBody
	GetTargets() []*string
	SetTotal(v int32) *ListEnterpriseAccelerateTargetsResponseBody
	GetTotal() *int32
}

type ListEnterpriseAccelerateTargetsResponseBody struct {
	// example:
	//
	// eap-7fed37a757a0de24
	EapId *string `json:"EapId,omitempty" xml:"EapId,omitempty"`
	// example:
	//
	// 529F755E-2E75-52EC-9C2E-6293FB8BF986
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Targets   []*string `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
	// example:
	//
	// 103
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListEnterpriseAccelerateTargetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseAccelerateTargetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnterpriseAccelerateTargetsResponseBody) GetEapId() *string {
	return s.EapId
}

func (s *ListEnterpriseAccelerateTargetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEnterpriseAccelerateTargetsResponseBody) GetTargets() []*string {
	return s.Targets
}

func (s *ListEnterpriseAccelerateTargetsResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListEnterpriseAccelerateTargetsResponseBody) SetEapId(v string) *ListEnterpriseAccelerateTargetsResponseBody {
	s.EapId = &v
	return s
}

func (s *ListEnterpriseAccelerateTargetsResponseBody) SetRequestId(v string) *ListEnterpriseAccelerateTargetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEnterpriseAccelerateTargetsResponseBody) SetTargets(v []*string) *ListEnterpriseAccelerateTargetsResponseBody {
	s.Targets = v
	return s
}

func (s *ListEnterpriseAccelerateTargetsResponseBody) SetTotal(v int32) *ListEnterpriseAccelerateTargetsResponseBody {
	s.Total = &v
	return s
}

func (s *ListEnterpriseAccelerateTargetsResponseBody) Validate() error {
	return dara.Validate(s)
}
