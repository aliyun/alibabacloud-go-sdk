// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveVpcAccessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApis(v *RemoveVpcAccessResponseBodyApis) *RemoveVpcAccessResponseBody
	GetApis() *RemoveVpcAccessResponseBodyApis
	SetRequestId(v string) *RemoveVpcAccessResponseBody
	GetRequestId() *string
}

type RemoveVpcAccessResponseBody struct {
	// API operations
	Apis *RemoveVpcAccessResponseBodyApis `json:"Apis,omitempty" xml:"Apis,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ015
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveVpcAccessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveVpcAccessResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveVpcAccessResponseBody) GetApis() *RemoveVpcAccessResponseBodyApis {
	return s.Apis
}

func (s *RemoveVpcAccessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveVpcAccessResponseBody) SetApis(v *RemoveVpcAccessResponseBodyApis) *RemoveVpcAccessResponseBody {
	s.Apis = v
	return s
}

func (s *RemoveVpcAccessResponseBody) SetRequestId(v string) *RemoveVpcAccessResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveVpcAccessResponseBody) Validate() error {
	return dara.Validate(s)
}

type RemoveVpcAccessResponseBodyApis struct {
	Api []*RemoveVpcAccessResponseBodyApisApi `json:"Api,omitempty" xml:"Api,omitempty" type:"Repeated"`
}

func (s RemoveVpcAccessResponseBodyApis) String() string {
	return dara.Prettify(s)
}

func (s RemoveVpcAccessResponseBodyApis) GoString() string {
	return s.String()
}

func (s *RemoveVpcAccessResponseBodyApis) GetApi() []*RemoveVpcAccessResponseBodyApisApi {
	return s.Api
}

func (s *RemoveVpcAccessResponseBodyApis) SetApi(v []*RemoveVpcAccessResponseBodyApisApi) *RemoveVpcAccessResponseBodyApis {
	s.Api = v
	return s
}

func (s *RemoveVpcAccessResponseBodyApis) Validate() error {
	return dara.Validate(s)
}

type RemoveVpcAccessResponseBodyApisApi struct {
	// API Id
	//
	// example:
	//
	// 551877242a4b4f3a84a56b7c3570e4a7
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The ID of the API group.
	//
	// example:
	//
	// 78d54ac4424d4b1792e33ca35637e8e4
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the runtime environment.
	//
	// example:
	//
	// d1e1ee28f9fb4b729db0ee8ca76ff0a5
	StageId *string `json:"StageId,omitempty" xml:"StageId,omitempty"`
}

func (s RemoveVpcAccessResponseBodyApisApi) String() string {
	return dara.Prettify(s)
}

func (s RemoveVpcAccessResponseBodyApisApi) GoString() string {
	return s.String()
}

func (s *RemoveVpcAccessResponseBodyApisApi) GetApiId() *string {
	return s.ApiId
}

func (s *RemoveVpcAccessResponseBodyApisApi) GetGroupId() *string {
	return s.GroupId
}

func (s *RemoveVpcAccessResponseBodyApisApi) GetStageId() *string {
	return s.StageId
}

func (s *RemoveVpcAccessResponseBodyApisApi) SetApiId(v string) *RemoveVpcAccessResponseBodyApisApi {
	s.ApiId = &v
	return s
}

func (s *RemoveVpcAccessResponseBodyApisApi) SetGroupId(v string) *RemoveVpcAccessResponseBodyApisApi {
	s.GroupId = &v
	return s
}

func (s *RemoveVpcAccessResponseBodyApisApi) SetStageId(v string) *RemoveVpcAccessResponseBodyApisApi {
	s.StageId = &v
	return s
}

func (s *RemoveVpcAccessResponseBodyApisApi) Validate() error {
	return dara.Validate(s)
}
