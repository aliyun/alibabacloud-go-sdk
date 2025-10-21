// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitPreviewResult interface {
	dara.Model
	String() string
	GoString() string
	SetQueryId(v string) *SubmitPreviewResult
	GetQueryId() *string
	SetSessionId(v string) *SubmitPreviewResult
	GetSessionId() *string
	SetTableSchemas(v []*TableSchema) *SubmitPreviewResult
	GetTableSchemas() []*TableSchema
}

type SubmitPreviewResult struct {
	QueryId      *string        `json:"queryId,omitempty" xml:"queryId,omitempty"`
	SessionId    *string        `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	TableSchemas []*TableSchema `json:"tableSchemas,omitempty" xml:"tableSchemas,omitempty" type:"Repeated"`
}

func (s SubmitPreviewResult) String() string {
	return dara.Prettify(s)
}

func (s SubmitPreviewResult) GoString() string {
	return s.String()
}

func (s *SubmitPreviewResult) GetQueryId() *string {
	return s.QueryId
}

func (s *SubmitPreviewResult) GetSessionId() *string {
	return s.SessionId
}

func (s *SubmitPreviewResult) GetTableSchemas() []*TableSchema {
	return s.TableSchemas
}

func (s *SubmitPreviewResult) SetQueryId(v string) *SubmitPreviewResult {
	s.QueryId = &v
	return s
}

func (s *SubmitPreviewResult) SetSessionId(v string) *SubmitPreviewResult {
	s.SessionId = &v
	return s
}

func (s *SubmitPreviewResult) SetTableSchemas(v []*TableSchema) *SubmitPreviewResult {
	s.TableSchemas = v
	return s
}

func (s *SubmitPreviewResult) Validate() error {
	if s.TableSchemas != nil {
		for _, item := range s.TableSchemas {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
