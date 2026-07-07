// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecureSkillIdentitiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIdentities(v []*string) *ListSecureSkillIdentitiesResponseBody
	GetIdentities() []*string
	SetRequestId(v string) *ListSecureSkillIdentitiesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListSecureSkillIdentitiesResponseBody
	GetTotalCount() *int64
}

type ListSecureSkillIdentitiesResponseBody struct {
	// The list of resource information.
	Identities []*string `json:"Identities,omitempty" xml:"Identities,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 51592A88-0F2C-55E6-AD2C-2AD9C10D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 20
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSecureSkillIdentitiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSecureSkillIdentitiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSecureSkillIdentitiesResponseBody) GetIdentities() []*string {
	return s.Identities
}

func (s *ListSecureSkillIdentitiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSecureSkillIdentitiesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListSecureSkillIdentitiesResponseBody) SetIdentities(v []*string) *ListSecureSkillIdentitiesResponseBody {
	s.Identities = v
	return s
}

func (s *ListSecureSkillIdentitiesResponseBody) SetRequestId(v string) *ListSecureSkillIdentitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSecureSkillIdentitiesResponseBody) SetTotalCount(v int64) *ListSecureSkillIdentitiesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSecureSkillIdentitiesResponseBody) Validate() error {
	return dara.Validate(s)
}
