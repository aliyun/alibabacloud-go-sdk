// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMultimodalSearchOutput interface {
	dara.Model
	String() string
	GoString() string
	SetImageItems(v []*UnifiedImageItem) *MultimodalSearchOutput
	GetImageItems() []*UnifiedImageItem
	SetQueryContext(v *MultimodalQueryContext) *MultimodalSearchOutput
	GetQueryContext() *MultimodalQueryContext
	SetRequestId(v string) *MultimodalSearchOutput
	GetRequestId() *string
	SetSearchInformation(v *SearchInformation) *MultimodalSearchOutput
	GetSearchInformation() *SearchInformation
}

type MultimodalSearchOutput struct {
	ImageItems        []*UnifiedImageItem     `json:"imageItems,omitempty" xml:"imageItems,omitempty" type:"Repeated"`
	QueryContext      *MultimodalQueryContext `json:"queryContext,omitempty" xml:"queryContext,omitempty"`
	RequestId         *string                 `json:"requestId,omitempty" xml:"requestId,omitempty"`
	SearchInformation *SearchInformation      `json:"searchInformation,omitempty" xml:"searchInformation,omitempty"`
}

func (s MultimodalSearchOutput) String() string {
	return dara.Prettify(s)
}

func (s MultimodalSearchOutput) GoString() string {
	return s.String()
}

func (s *MultimodalSearchOutput) GetImageItems() []*UnifiedImageItem {
	return s.ImageItems
}

func (s *MultimodalSearchOutput) GetQueryContext() *MultimodalQueryContext {
	return s.QueryContext
}

func (s *MultimodalSearchOutput) GetRequestId() *string {
	return s.RequestId
}

func (s *MultimodalSearchOutput) GetSearchInformation() *SearchInformation {
	return s.SearchInformation
}

func (s *MultimodalSearchOutput) SetImageItems(v []*UnifiedImageItem) *MultimodalSearchOutput {
	s.ImageItems = v
	return s
}

func (s *MultimodalSearchOutput) SetQueryContext(v *MultimodalQueryContext) *MultimodalSearchOutput {
	s.QueryContext = v
	return s
}

func (s *MultimodalSearchOutput) SetRequestId(v string) *MultimodalSearchOutput {
	s.RequestId = &v
	return s
}

func (s *MultimodalSearchOutput) SetSearchInformation(v *SearchInformation) *MultimodalSearchOutput {
	s.SearchInformation = v
	return s
}

func (s *MultimodalSearchOutput) Validate() error {
	if s.ImageItems != nil {
		for _, item := range s.ImageItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.QueryContext != nil {
		if err := s.QueryContext.Validate(); err != nil {
			return err
		}
	}
	if s.SearchInformation != nil {
		if err := s.SearchInformation.Validate(); err != nil {
			return err
		}
	}
	return nil
}
