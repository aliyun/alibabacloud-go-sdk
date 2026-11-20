// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTextbookAssistantSceneDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthToken(v string) *ListTextbookAssistantSceneDetailsRequest
	GetAuthToken() *string
	SetSceneIdList(v []*string) *ListTextbookAssistantSceneDetailsRequest
	GetSceneIdList() []*string
}

type ListTextbookAssistantSceneDetailsRequest struct {
	// The authorization code for API calls. You can obtain this code by calling the API operation to retrieve the authorization code for the English textbook-style AI teacher feature.
	//
	// example:
	//
	// tc_e6dc70c890866f4028ca685b6fa29874
	AuthToken *string `json:"authToken,omitempty" xml:"authToken,omitempty"`
	// The list of scene IDs.
	SceneIdList []*string `json:"sceneIdList,omitempty" xml:"sceneIdList,omitempty" type:"Repeated"`
}

func (s ListTextbookAssistantSceneDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantSceneDetailsRequest) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantSceneDetailsRequest) GetAuthToken() *string {
	return s.AuthToken
}

func (s *ListTextbookAssistantSceneDetailsRequest) GetSceneIdList() []*string {
	return s.SceneIdList
}

func (s *ListTextbookAssistantSceneDetailsRequest) SetAuthToken(v string) *ListTextbookAssistantSceneDetailsRequest {
	s.AuthToken = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsRequest) SetSceneIdList(v []*string) *ListTextbookAssistantSceneDetailsRequest {
	s.SceneIdList = v
	return s
}

func (s *ListTextbookAssistantSceneDetailsRequest) Validate() error {
	return dara.Validate(s)
}
