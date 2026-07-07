// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillAuthedIdentitiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIdentities(v []*ListSkillAuthedIdentitiesResponseBodyIdentities) *ListSkillAuthedIdentitiesResponseBody
	GetIdentities() []*ListSkillAuthedIdentitiesResponseBodyIdentities
	SetRequestId(v string) *ListSkillAuthedIdentitiesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListSkillAuthedIdentitiesResponseBody
	GetTotalCount() *int64
}

type ListSkillAuthedIdentitiesResponseBody struct {
	// The list of authorized objects.
	Identities []*ListSkillAuthedIdentitiesResponseBodyIdentities `json:"Identities,omitempty" xml:"Identities,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 5CC5E450-FC43-4F5B-B540-9964BD*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 20
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSkillAuthedIdentitiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSkillAuthedIdentitiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSkillAuthedIdentitiesResponseBody) GetIdentities() []*ListSkillAuthedIdentitiesResponseBodyIdentities {
	return s.Identities
}

func (s *ListSkillAuthedIdentitiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSkillAuthedIdentitiesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListSkillAuthedIdentitiesResponseBody) SetIdentities(v []*ListSkillAuthedIdentitiesResponseBodyIdentities) *ListSkillAuthedIdentitiesResponseBody {
	s.Identities = v
	return s
}

func (s *ListSkillAuthedIdentitiesResponseBody) SetRequestId(v string) *ListSkillAuthedIdentitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSkillAuthedIdentitiesResponseBody) SetTotalCount(v int64) *ListSkillAuthedIdentitiesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSkillAuthedIdentitiesResponseBody) Validate() error {
	if s.Identities != nil {
		for _, item := range s.Identities {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSkillAuthedIdentitiesResponseBodyIdentities struct {
	// Indicates whether automatic installation is enabled. Valid values:
	//
	// - true: Automatic installation is enabled.
	//
	// - false: Automatic installation is disabled.
	//
	// example:
	//
	// true
	AutoInstall *bool `json:"AutoInstall,omitempty" xml:"AutoInstall,omitempty"`
	// The ID of the authorized object.
	//
	// example:
	//
	// ecd-b9ej3xiok4tjbgf9x****
	IdentityId *string `json:"IdentityId,omitempty" xml:"IdentityId,omitempty"`
}

func (s ListSkillAuthedIdentitiesResponseBodyIdentities) String() string {
	return dara.Prettify(s)
}

func (s ListSkillAuthedIdentitiesResponseBodyIdentities) GoString() string {
	return s.String()
}

func (s *ListSkillAuthedIdentitiesResponseBodyIdentities) GetAutoInstall() *bool {
	return s.AutoInstall
}

func (s *ListSkillAuthedIdentitiesResponseBodyIdentities) GetIdentityId() *string {
	return s.IdentityId
}

func (s *ListSkillAuthedIdentitiesResponseBodyIdentities) SetAutoInstall(v bool) *ListSkillAuthedIdentitiesResponseBodyIdentities {
	s.AutoInstall = &v
	return s
}

func (s *ListSkillAuthedIdentitiesResponseBodyIdentities) SetIdentityId(v string) *ListSkillAuthedIdentitiesResponseBodyIdentities {
	s.IdentityId = &v
	return s
}

func (s *ListSkillAuthedIdentitiesResponseBodyIdentities) Validate() error {
	return dara.Validate(s)
}
