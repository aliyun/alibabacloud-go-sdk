// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudNotePhrasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *CreateCloudNotePhrasesResponseBody
	GetId() *string
	SetRequestId(v string) *CreateCloudNotePhrasesResponseBody
	GetRequestId() *string
}

type CreateCloudNotePhrasesResponseBody struct {
	// example:
	//
	// 21088b2617489486958531017d0b19
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCloudNotePhrasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudNotePhrasesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCloudNotePhrasesResponseBody) GetId() *string {
	return s.Id
}

func (s *CreateCloudNotePhrasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCloudNotePhrasesResponseBody) SetId(v string) *CreateCloudNotePhrasesResponseBody {
	s.Id = &v
	return s
}

func (s *CreateCloudNotePhrasesResponseBody) SetRequestId(v string) *CreateCloudNotePhrasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCloudNotePhrasesResponseBody) Validate() error {
	return dara.Validate(s)
}
