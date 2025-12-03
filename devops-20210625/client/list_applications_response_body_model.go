// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListApplicationsResponseBodyData) *ListApplicationsResponseBody
	GetData() []*ListApplicationsResponseBodyData
	SetNextToken(v string) *ListApplicationsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListApplicationsResponseBody
	GetRequestId() *string
}

type ListApplicationsResponseBody struct {
	Data []*ListApplicationsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// vxc2341gfssad12
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// FC93CE1A-8D7A-13A9-8306-7465DE2E5C0F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListApplicationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBody) GetData() []*ListApplicationsResponseBodyData {
	return s.Data
}

func (s *ListApplicationsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListApplicationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationsResponseBody) SetData(v []*ListApplicationsResponseBodyData) *ListApplicationsResponseBody {
	s.Data = v
	return s
}

func (s *ListApplicationsResponseBody) SetNextToken(v string) *ListApplicationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListApplicationsResponseBody) SetRequestId(v string) *ListApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApplicationsResponseBodyData struct {
	// example:
	//
	// 1332695887xxxxxx
	CreatorAccountId *string `json:"creatorAccountId,omitempty" xml:"creatorAccountId,omitempty"`
	// example:
	//
	// 应用描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 2024-01-01T00:00:00.000+00:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// testApp
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListApplicationsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBodyData) GetCreatorAccountId() *string {
	return s.CreatorAccountId
}

func (s *ListApplicationsResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListApplicationsResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListApplicationsResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListApplicationsResponseBodyData) SetCreatorAccountId(v string) *ListApplicationsResponseBodyData {
	s.CreatorAccountId = &v
	return s
}

func (s *ListApplicationsResponseBodyData) SetDescription(v string) *ListApplicationsResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListApplicationsResponseBodyData) SetGmtCreate(v string) *ListApplicationsResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListApplicationsResponseBodyData) SetName(v string) *ListApplicationsResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListApplicationsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
