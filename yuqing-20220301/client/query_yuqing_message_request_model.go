// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryYuqingMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryYuqingMessageRequest
	GetRequestId() *string
	SetSearchCondition(v *SearchCondition) *QueryYuqingMessageRequest
	GetSearchCondition() *SearchCondition
	SetTeamHashId(v string) *QueryYuqingMessageRequest
	GetTeamHashId() *string
}

type QueryYuqingMessageRequest struct {
	// example:
	//
	// 5645a6c9-7d21-4926-a410-db9a1af85faa
	RequestId       *string          `json:"requestId,omitempty" xml:"requestId,omitempty"`
	SearchCondition *SearchCondition `json:"searchCondition,omitempty" xml:"searchCondition,omitempty"`
	// example:
	//
	// xxxx43434dsdsd
	TeamHashId *string `json:"teamHashId,omitempty" xml:"teamHashId,omitempty"`
}

func (s QueryYuqingMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryYuqingMessageRequest) GoString() string {
	return s.String()
}

func (s *QueryYuqingMessageRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryYuqingMessageRequest) GetSearchCondition() *SearchCondition {
	return s.SearchCondition
}

func (s *QueryYuqingMessageRequest) GetTeamHashId() *string {
	return s.TeamHashId
}

func (s *QueryYuqingMessageRequest) SetRequestId(v string) *QueryYuqingMessageRequest {
	s.RequestId = &v
	return s
}

func (s *QueryYuqingMessageRequest) SetSearchCondition(v *SearchCondition) *QueryYuqingMessageRequest {
	s.SearchCondition = v
	return s
}

func (s *QueryYuqingMessageRequest) SetTeamHashId(v string) *QueryYuqingMessageRequest {
	s.TeamHashId = &v
	return s
}

func (s *QueryYuqingMessageRequest) Validate() error {
	if s.SearchCondition != nil {
		if err := s.SearchCondition.Validate(); err != nil {
			return err
		}
	}
	return nil
}
