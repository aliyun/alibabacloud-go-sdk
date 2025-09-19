// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateMaliciousNoteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageMaliciousFileList(v []*BatchCreateMaliciousNoteRequestImageMaliciousFileList) *BatchCreateMaliciousNoteRequest
	GetImageMaliciousFileList() []*BatchCreateMaliciousNoteRequestImageMaliciousFileList
}

type BatchCreateMaliciousNoteRequest struct {
	// The batches.
	ImageMaliciousFileList []*BatchCreateMaliciousNoteRequestImageMaliciousFileList `json:"ImageMaliciousFileList,omitempty" xml:"ImageMaliciousFileList,omitempty" type:"Repeated"`
}

func (s BatchCreateMaliciousNoteRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateMaliciousNoteRequest) GoString() string {
	return s.String()
}

func (s *BatchCreateMaliciousNoteRequest) GetImageMaliciousFileList() []*BatchCreateMaliciousNoteRequestImageMaliciousFileList {
	return s.ImageMaliciousFileList
}

func (s *BatchCreateMaliciousNoteRequest) SetImageMaliciousFileList(v []*BatchCreateMaliciousNoteRequestImageMaliciousFileList) *BatchCreateMaliciousNoteRequest {
	s.ImageMaliciousFileList = v
	return s
}

func (s *BatchCreateMaliciousNoteRequest) Validate() error {
	return dara.Validate(s)
}

type BatchCreateMaliciousNoteRequestImageMaliciousFileList struct {
	// The ID of the alert.
	//
	// >  You can call the [DescribeSuspEvents](~~DescribeSuspEvents~~) operation to query the alert IDs.
	//
	// example:
	//
	// 1
	EventId *int64 `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The description.
	//
	// example:
	//
	// Malware sample
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
}

func (s BatchCreateMaliciousNoteRequestImageMaliciousFileList) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateMaliciousNoteRequestImageMaliciousFileList) GoString() string {
	return s.String()
}

func (s *BatchCreateMaliciousNoteRequestImageMaliciousFileList) GetEventId() *int64 {
	return s.EventId
}

func (s *BatchCreateMaliciousNoteRequestImageMaliciousFileList) GetNote() *string {
	return s.Note
}

func (s *BatchCreateMaliciousNoteRequestImageMaliciousFileList) SetEventId(v int64) *BatchCreateMaliciousNoteRequestImageMaliciousFileList {
	s.EventId = &v
	return s
}

func (s *BatchCreateMaliciousNoteRequestImageMaliciousFileList) SetNote(v string) *BatchCreateMaliciousNoteRequestImageMaliciousFileList {
	s.Note = &v
	return s
}

func (s *BatchCreateMaliciousNoteRequestImageMaliciousFileList) Validate() error {
	return dara.Validate(s)
}
