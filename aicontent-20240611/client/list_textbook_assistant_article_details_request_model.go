// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTextbookAssistantArticleDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArticleIdList(v []*string) *ListTextbookAssistantArticleDetailsRequest
	GetArticleIdList() []*string
	SetAuthToken(v string) *ListTextbookAssistantArticleDetailsRequest
	GetAuthToken() *string
}

type ListTextbookAssistantArticleDetailsRequest struct {
	ArticleIdList []*string `json:"articleIdList,omitempty" xml:"articleIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// tc_e6dc70c890866f4028ca685b6fa29874
	AuthToken *string `json:"authToken,omitempty" xml:"authToken,omitempty"`
}

func (s ListTextbookAssistantArticleDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantArticleDetailsRequest) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantArticleDetailsRequest) GetArticleIdList() []*string {
	return s.ArticleIdList
}

func (s *ListTextbookAssistantArticleDetailsRequest) GetAuthToken() *string {
	return s.AuthToken
}

func (s *ListTextbookAssistantArticleDetailsRequest) SetArticleIdList(v []*string) *ListTextbookAssistantArticleDetailsRequest {
	s.ArticleIdList = v
	return s
}

func (s *ListTextbookAssistantArticleDetailsRequest) SetAuthToken(v string) *ListTextbookAssistantArticleDetailsRequest {
	s.AuthToken = &v
	return s
}

func (s *ListTextbookAssistantArticleDetailsRequest) Validate() error {
	return dara.Validate(s)
}
