// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryStoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *QueryStoriesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *QueryStoriesResponseBody
	GetRequestId() *string
	SetStories(v []*Story) *QueryStoriesResponseBody
	GetStories() []*Story
}

type QueryStoriesResponseBody struct {
	// The pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// MTIzNDU2Nzg6aW1tdGVzdDpleGFtcGxlYnVja2V0OmRhdGFzZXQwMDE6b3NzOi8vZXhhbXBsZWJ1Y2tldC9zYW1wbGVvYmplY3Qx****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2C5C1E0F-D8B8-4DA0-8127-EC32C771****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The stories.
	Stories []*Story `json:"Stories,omitempty" xml:"Stories,omitempty" type:"Repeated"`
}

func (s QueryStoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryStoriesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryStoriesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryStoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryStoriesResponseBody) GetStories() []*Story {
	return s.Stories
}

func (s *QueryStoriesResponseBody) SetNextToken(v string) *QueryStoriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryStoriesResponseBody) SetRequestId(v string) *QueryStoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryStoriesResponseBody) SetStories(v []*Story) *QueryStoriesResponseBody {
	s.Stories = v
	return s
}

func (s *QueryStoriesResponseBody) Validate() error {
	if s.Stories != nil {
		for _, item := range s.Stories {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
