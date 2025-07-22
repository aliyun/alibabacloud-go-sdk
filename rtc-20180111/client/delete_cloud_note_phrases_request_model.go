// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudNotePhrasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteCloudNotePhrasesRequest
	GetAppId() *string
	SetPhrase(v *DeleteCloudNotePhrasesRequestPhrase) *DeleteCloudNotePhrasesRequest
	GetPhrase() *DeleteCloudNotePhrasesRequestPhrase
}

type DeleteCloudNotePhrasesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	Phrase *DeleteCloudNotePhrasesRequestPhrase `json:"Phrase,omitempty" xml:"Phrase,omitempty" type:"Struct"`
}

func (s DeleteCloudNotePhrasesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudNotePhrasesRequest) GoString() string {
	return s.String()
}

func (s *DeleteCloudNotePhrasesRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteCloudNotePhrasesRequest) GetPhrase() *DeleteCloudNotePhrasesRequestPhrase {
	return s.Phrase
}

func (s *DeleteCloudNotePhrasesRequest) SetAppId(v string) *DeleteCloudNotePhrasesRequest {
	s.AppId = &v
	return s
}

func (s *DeleteCloudNotePhrasesRequest) SetPhrase(v *DeleteCloudNotePhrasesRequestPhrase) *DeleteCloudNotePhrasesRequest {
	s.Phrase = v
	return s
}

func (s *DeleteCloudNotePhrasesRequest) Validate() error {
	return dara.Validate(s)
}

type DeleteCloudNotePhrasesRequestPhrase struct {
	// This parameter is required.
	//
	// example:
	//
	// 1qasw23ezcsrfsawq
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteCloudNotePhrasesRequestPhrase) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudNotePhrasesRequestPhrase) GoString() string {
	return s.String()
}

func (s *DeleteCloudNotePhrasesRequestPhrase) GetId() *string {
	return s.Id
}

func (s *DeleteCloudNotePhrasesRequestPhrase) SetId(v string) *DeleteCloudNotePhrasesRequestPhrase {
	s.Id = &v
	return s
}

func (s *DeleteCloudNotePhrasesRequestPhrase) Validate() error {
	return dara.Validate(s)
}
