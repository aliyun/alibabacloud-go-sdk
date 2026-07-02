// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPaperDescription interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v []*Summary) *PaperDescription
	GetDescription() []*Summary
	SetTitleID(v []*string) *PaperDescription
	GetTitleID() []*string
}

type PaperDescription struct {
	// The guide result.
	Description []*Summary `json:"Description,omitempty" xml:"Description,omitempty" type:"Repeated"`
	// The section heading included in the guide result.
	TitleID []*string `json:"TitleID,omitempty" xml:"TitleID,omitempty" type:"Repeated"`
}

func (s PaperDescription) String() string {
	return dara.Prettify(s)
}

func (s PaperDescription) GoString() string {
	return s.String()
}

func (s *PaperDescription) GetDescription() []*Summary {
	return s.Description
}

func (s *PaperDescription) GetTitleID() []*string {
	return s.TitleID
}

func (s *PaperDescription) SetDescription(v []*Summary) *PaperDescription {
	s.Description = v
	return s
}

func (s *PaperDescription) SetTitleID(v []*string) *PaperDescription {
	s.TitleID = v
	return s
}

func (s *PaperDescription) Validate() error {
	if s.Description != nil {
		for _, item := range s.Description {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
