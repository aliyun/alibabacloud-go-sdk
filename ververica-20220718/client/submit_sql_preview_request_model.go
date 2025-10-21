// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSqlPreviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *SqlStatementWithContext) *SubmitSqlPreviewRequest
	GetBody() *SqlStatementWithContext
	SetSessionClusterName(v string) *SubmitSqlPreviewRequest
	GetSessionClusterName() *string
}

type SubmitSqlPreviewRequest struct {
	Body *SqlStatementWithContext `json:"body,omitempty" xml:"body,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-session
	SessionClusterName *string `json:"sessionClusterName,omitempty" xml:"sessionClusterName,omitempty"`
}

func (s SubmitSqlPreviewRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitSqlPreviewRequest) GoString() string {
	return s.String()
}

func (s *SubmitSqlPreviewRequest) GetBody() *SqlStatementWithContext {
	return s.Body
}

func (s *SubmitSqlPreviewRequest) GetSessionClusterName() *string {
	return s.SessionClusterName
}

func (s *SubmitSqlPreviewRequest) SetBody(v *SqlStatementWithContext) *SubmitSqlPreviewRequest {
	s.Body = v
	return s
}

func (s *SubmitSqlPreviewRequest) SetSessionClusterName(v string) *SubmitSqlPreviewRequest {
	s.SessionClusterName = &v
	return s
}

func (s *SubmitSqlPreviewRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
