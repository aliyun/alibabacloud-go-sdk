// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImportJobList(v *ListJobResp) *ListJobResponseBody
	GetImportJobList() *ListJobResp
}

type ListJobResponseBody struct {
	// The queried migration tasks.
	ImportJobList *ListJobResp `json:"ImportJobList,omitempty" xml:"ImportJobList,omitempty"`
}

func (s ListJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobResponseBody) GetImportJobList() *ListJobResp {
	return s.ImportJobList
}

func (s *ListJobResponseBody) SetImportJobList(v *ListJobResp) *ListJobResponseBody {
	s.ImportJobList = v
	return s
}

func (s *ListJobResponseBody) Validate() error {
	if s.ImportJobList != nil {
		if err := s.ImportJobList.Validate(); err != nil {
			return err
		}
	}
	return nil
}
