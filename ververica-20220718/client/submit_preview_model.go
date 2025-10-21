// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitPreview interface {
	dara.Model
	String() string
	GoString() string
	SetQueryName(v string) *SubmitPreview
	GetQueryName() *string
	SetSessionId(v string) *SubmitPreview
	GetSessionId() *string
	SetTableSchemas(v []*TableSchema) *SubmitPreview
	GetTableSchemas() []*TableSchema
}

type SubmitPreview struct {
	QueryName    *string        `json:"queryName,omitempty" xml:"queryName,omitempty"`
	SessionId    *string        `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	TableSchemas []*TableSchema `json:"tableSchemas,omitempty" xml:"tableSchemas,omitempty" type:"Repeated"`
}

func (s SubmitPreview) String() string {
	return dara.Prettify(s)
}

func (s SubmitPreview) GoString() string {
	return s.String()
}

func (s *SubmitPreview) GetQueryName() *string {
	return s.QueryName
}

func (s *SubmitPreview) GetSessionId() *string {
	return s.SessionId
}

func (s *SubmitPreview) GetTableSchemas() []*TableSchema {
	return s.TableSchemas
}

func (s *SubmitPreview) SetQueryName(v string) *SubmitPreview {
	s.QueryName = &v
	return s
}

func (s *SubmitPreview) SetSessionId(v string) *SubmitPreview {
	s.SessionId = &v
	return s
}

func (s *SubmitPreview) SetTableSchemas(v []*TableSchema) *SubmitPreview {
	s.TableSchemas = v
	return s
}

func (s *SubmitPreview) Validate() error {
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
