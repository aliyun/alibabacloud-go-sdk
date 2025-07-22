// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCloudNotePhrasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *ModifyCloudNotePhrasesResponseBody
	GetId() *string
	SetRequestId(v string) *ModifyCloudNotePhrasesResponseBody
	GetRequestId() *string
}

type ModifyCloudNotePhrasesResponseBody struct {
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

func (s ModifyCloudNotePhrasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudNotePhrasesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCloudNotePhrasesResponseBody) GetId() *string {
	return s.Id
}

func (s *ModifyCloudNotePhrasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCloudNotePhrasesResponseBody) SetId(v string) *ModifyCloudNotePhrasesResponseBody {
	s.Id = &v
	return s
}

func (s *ModifyCloudNotePhrasesResponseBody) SetRequestId(v string) *ModifyCloudNotePhrasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCloudNotePhrasesResponseBody) Validate() error {
	return dara.Validate(s)
}
