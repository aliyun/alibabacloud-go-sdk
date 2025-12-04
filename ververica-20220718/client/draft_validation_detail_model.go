// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDraftValidationDetail interface {
	dara.Model
	String() string
	GoString() string
	SetDraftMetaInfoErrorDetails(v []*DraftMetaInfoErrorDetail) *DraftValidationDetail
	GetDraftMetaInfoErrorDetails() []*DraftMetaInfoErrorDetail
	SetSqlErrorDetail(v *ValidationErrorDetails) *DraftValidationDetail
	GetSqlErrorDetail() *ValidationErrorDetails
	SetSqlValidationResult(v string) *DraftValidationDetail
	GetSqlValidationResult() *string
}

type DraftValidationDetail struct {
	DraftMetaInfoErrorDetails []*DraftMetaInfoErrorDetail `json:"draftMetaInfoErrorDetails,omitempty" xml:"draftMetaInfoErrorDetails,omitempty" type:"Repeated"`
	SqlErrorDetail            *ValidationErrorDetails     `json:"sqlErrorDetail,omitempty" xml:"sqlErrorDetail,omitempty"`
	SqlValidationResult       *string                     `json:"sqlValidationResult,omitempty" xml:"sqlValidationResult,omitempty"`
}

func (s DraftValidationDetail) String() string {
	return dara.Prettify(s)
}

func (s DraftValidationDetail) GoString() string {
	return s.String()
}

func (s *DraftValidationDetail) GetDraftMetaInfoErrorDetails() []*DraftMetaInfoErrorDetail {
	return s.DraftMetaInfoErrorDetails
}

func (s *DraftValidationDetail) GetSqlErrorDetail() *ValidationErrorDetails {
	return s.SqlErrorDetail
}

func (s *DraftValidationDetail) GetSqlValidationResult() *string {
	return s.SqlValidationResult
}

func (s *DraftValidationDetail) SetDraftMetaInfoErrorDetails(v []*DraftMetaInfoErrorDetail) *DraftValidationDetail {
	s.DraftMetaInfoErrorDetails = v
	return s
}

func (s *DraftValidationDetail) SetSqlErrorDetail(v *ValidationErrorDetails) *DraftValidationDetail {
	s.SqlErrorDetail = v
	return s
}

func (s *DraftValidationDetail) SetSqlValidationResult(v string) *DraftValidationDetail {
	s.SqlValidationResult = &v
	return s
}

func (s *DraftValidationDetail) Validate() error {
	if s.DraftMetaInfoErrorDetails != nil {
		for _, item := range s.DraftMetaInfoErrorDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SqlErrorDetail != nil {
		if err := s.SqlErrorDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}
