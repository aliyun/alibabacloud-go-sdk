// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchSemanticViewsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*SemanticViewModel) *SearchSemanticViewsResponseBody
	GetData() []*SemanticViewModel
	SetRequestId(v string) *SearchSemanticViewsResponseBody
	GetRequestId() *string
}

type SearchSemanticViewsResponseBody struct {
	// A list of semantic view objects.
	Data []*SemanticViewModel `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SearchSemanticViewsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchSemanticViewsResponseBody) GoString() string {
	return s.String()
}

func (s *SearchSemanticViewsResponseBody) GetData() []*SemanticViewModel {
	return s.Data
}

func (s *SearchSemanticViewsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchSemanticViewsResponseBody) SetData(v []*SemanticViewModel) *SearchSemanticViewsResponseBody {
	s.Data = v
	return s
}

func (s *SearchSemanticViewsResponseBody) SetRequestId(v string) *SearchSemanticViewsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchSemanticViewsResponseBody) Validate() error {
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
