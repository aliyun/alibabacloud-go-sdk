// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDpiGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDpiGroup(v []*ListDpiGroupsResponseBodyDpiGroup) *ListDpiGroupsResponseBody
	GetDpiGroup() []*ListDpiGroupsResponseBodyDpiGroup
	SetNextToken(v string) *ListDpiGroupsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDpiGroupsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListDpiGroupsResponseBody
	GetTotalCount() *int32
}

type ListDpiGroupsResponseBody struct {
	// The information about the application group.
	DpiGroup []*ListDpiGroupsResponseBodyDpiGroup `json:"DpiGroup,omitempty" xml:"DpiGroup,omitempty" type:"Repeated"`
	// The token returned for the next query.
	//
	// example:
	//
	// FFPSpX59Ebw****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// EC184A86-3C93-49D6-BB34-6C193E14D37F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 22
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDpiGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDpiGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDpiGroupsResponseBody) GetDpiGroup() []*ListDpiGroupsResponseBodyDpiGroup {
	return s.DpiGroup
}

func (s *ListDpiGroupsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDpiGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDpiGroupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDpiGroupsResponseBody) SetDpiGroup(v []*ListDpiGroupsResponseBodyDpiGroup) *ListDpiGroupsResponseBody {
	s.DpiGroup = v
	return s
}

func (s *ListDpiGroupsResponseBody) SetNextToken(v string) *ListDpiGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDpiGroupsResponseBody) SetRequestId(v string) *ListDpiGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDpiGroupsResponseBody) SetTotalCount(v int32) *ListDpiGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDpiGroupsResponseBody) Validate() error {
	if s.DpiGroup != nil {
		for _, item := range s.DpiGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDpiGroupsResponseBodyDpiGroup struct {
	// The ID of the application group.
	//
	// example:
	//
	// 1
	DpiGroupId *string `json:"DpiGroupId,omitempty" xml:"DpiGroupId,omitempty"`
	// The name of the application group.
	//
	// example:
	//
	// P2P
	DpiGroupName *string `json:"DpiGroupName,omitempty" xml:"DpiGroupName,omitempty"`
	// The earliest version of engine that supports the application group.
	//
	// example:
	//
	// 0-0.0.1
	MinEngineVersion *string `json:"MinEngineVersion,omitempty" xml:"MinEngineVersion,omitempty"`
	// The earliest version of signature database that supports the application group.
	//
	// example:
	//
	// 20201117_1_0-0.0.1
	MinSignatureDbVersion *string `json:"MinSignatureDbVersion,omitempty" xml:"MinSignatureDbVersion,omitempty"`
}

func (s ListDpiGroupsResponseBodyDpiGroup) String() string {
	return dara.Prettify(s)
}

func (s ListDpiGroupsResponseBodyDpiGroup) GoString() string {
	return s.String()
}

func (s *ListDpiGroupsResponseBodyDpiGroup) GetDpiGroupId() *string {
	return s.DpiGroupId
}

func (s *ListDpiGroupsResponseBodyDpiGroup) GetDpiGroupName() *string {
	return s.DpiGroupName
}

func (s *ListDpiGroupsResponseBodyDpiGroup) GetMinEngineVersion() *string {
	return s.MinEngineVersion
}

func (s *ListDpiGroupsResponseBodyDpiGroup) GetMinSignatureDbVersion() *string {
	return s.MinSignatureDbVersion
}

func (s *ListDpiGroupsResponseBodyDpiGroup) SetDpiGroupId(v string) *ListDpiGroupsResponseBodyDpiGroup {
	s.DpiGroupId = &v
	return s
}

func (s *ListDpiGroupsResponseBodyDpiGroup) SetDpiGroupName(v string) *ListDpiGroupsResponseBodyDpiGroup {
	s.DpiGroupName = &v
	return s
}

func (s *ListDpiGroupsResponseBodyDpiGroup) SetMinEngineVersion(v string) *ListDpiGroupsResponseBodyDpiGroup {
	s.MinEngineVersion = &v
	return s
}

func (s *ListDpiGroupsResponseBodyDpiGroup) SetMinSignatureDbVersion(v string) *ListDpiGroupsResponseBodyDpiGroup {
	s.MinSignatureDbVersion = &v
	return s
}

func (s *ListDpiGroupsResponseBodyDpiGroup) Validate() error {
	return dara.Validate(s)
}
