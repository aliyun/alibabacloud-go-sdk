// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKernelReleaseNotesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetReleaseNotes(v *DescribeKernelReleaseNotesResponseBodyReleaseNotes) *DescribeKernelReleaseNotesResponseBody
	GetReleaseNotes() *DescribeKernelReleaseNotesResponseBodyReleaseNotes
	SetRequestId(v string) *DescribeKernelReleaseNotesResponseBody
	GetRequestId() *string
}

type DescribeKernelReleaseNotesResponseBody struct {
	// The list of the version release notes.
	ReleaseNotes *DescribeKernelReleaseNotesResponseBodyReleaseNotes `json:"ReleaseNotes,omitempty" xml:"ReleaseNotes,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F01D4DDA-CB72-4083-B399-AF4642294FE6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeKernelReleaseNotesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeKernelReleaseNotesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeKernelReleaseNotesResponseBody) GetReleaseNotes() *DescribeKernelReleaseNotesResponseBodyReleaseNotes {
	return s.ReleaseNotes
}

func (s *DescribeKernelReleaseNotesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeKernelReleaseNotesResponseBody) SetReleaseNotes(v *DescribeKernelReleaseNotesResponseBodyReleaseNotes) *DescribeKernelReleaseNotesResponseBody {
	s.ReleaseNotes = v
	return s
}

func (s *DescribeKernelReleaseNotesResponseBody) SetRequestId(v string) *DescribeKernelReleaseNotesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeKernelReleaseNotesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeKernelReleaseNotesResponseBodyReleaseNotes struct {
	ReleaseNote []*DescribeKernelReleaseNotesResponseBodyReleaseNotesReleaseNote `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty" type:"Repeated"`
}

func (s DescribeKernelReleaseNotesResponseBodyReleaseNotes) String() string {
	return dara.Prettify(s)
}

func (s DescribeKernelReleaseNotesResponseBodyReleaseNotes) GoString() string {
	return s.String()
}

func (s *DescribeKernelReleaseNotesResponseBodyReleaseNotes) GetReleaseNote() []*DescribeKernelReleaseNotesResponseBodyReleaseNotesReleaseNote {
	return s.ReleaseNote
}

func (s *DescribeKernelReleaseNotesResponseBodyReleaseNotes) SetReleaseNote(v []*DescribeKernelReleaseNotesResponseBodyReleaseNotesReleaseNote) *DescribeKernelReleaseNotesResponseBodyReleaseNotes {
	s.ReleaseNote = v
	return s
}

func (s *DescribeKernelReleaseNotesResponseBodyReleaseNotes) Validate() error {
	return dara.Validate(s)
}

type DescribeKernelReleaseNotesResponseBodyReleaseNotesReleaseNote struct {
	// The version number.
	//
	// example:
	//
	// mongodb_20180619_0.4.9
	KernelVersion *string `json:"KernelVersion,omitempty" xml:"KernelVersion,omitempty"`
	// The release notes.
	//
	// example:
	//
	// test release note.
	ReleaseNote *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
}

func (s DescribeKernelReleaseNotesResponseBodyReleaseNotesReleaseNote) String() string {
	return dara.Prettify(s)
}

func (s DescribeKernelReleaseNotesResponseBodyReleaseNotesReleaseNote) GoString() string {
	return s.String()
}

func (s *DescribeKernelReleaseNotesResponseBodyReleaseNotesReleaseNote) GetKernelVersion() *string {
	return s.KernelVersion
}

func (s *DescribeKernelReleaseNotesResponseBodyReleaseNotesReleaseNote) GetReleaseNote() *string {
	return s.ReleaseNote
}

func (s *DescribeKernelReleaseNotesResponseBodyReleaseNotesReleaseNote) SetKernelVersion(v string) *DescribeKernelReleaseNotesResponseBodyReleaseNotesReleaseNote {
	s.KernelVersion = &v
	return s
}

func (s *DescribeKernelReleaseNotesResponseBodyReleaseNotesReleaseNote) SetReleaseNote(v string) *DescribeKernelReleaseNotesResponseBodyReleaseNotesReleaseNote {
	s.ReleaseNote = &v
	return s
}

func (s *DescribeKernelReleaseNotesResponseBodyReleaseNotesReleaseNote) Validate() error {
	return dara.Validate(s)
}
