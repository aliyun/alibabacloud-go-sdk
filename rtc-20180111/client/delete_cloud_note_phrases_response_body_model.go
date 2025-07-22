// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudNotePhrasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteCloudNotePhrasesResponseBody
	GetId() *string
	SetRequestId(v string) *DeleteCloudNotePhrasesResponseBody
	GetRequestId() *string
}

type DeleteCloudNotePhrasesResponseBody struct {
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

func (s DeleteCloudNotePhrasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudNotePhrasesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCloudNotePhrasesResponseBody) GetId() *string {
	return s.Id
}

func (s *DeleteCloudNotePhrasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCloudNotePhrasesResponseBody) SetId(v string) *DeleteCloudNotePhrasesResponseBody {
	s.Id = &v
	return s
}

func (s *DeleteCloudNotePhrasesResponseBody) SetRequestId(v string) *DeleteCloudNotePhrasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCloudNotePhrasesResponseBody) Validate() error {
	return dara.Validate(s)
}
