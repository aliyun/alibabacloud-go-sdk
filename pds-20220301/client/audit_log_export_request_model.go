// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuditLogExportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *AuditLogExportRequest
	GetFileName() *string
	SetLanguage(v string) *AuditLogExportRequest
	GetLanguage() *string
	SetOrderBy(v string) *AuditLogExportRequest
	GetOrderBy() *string
	SetQuery(v string) *AuditLogExportRequest
	GetQuery() *string
}

type AuditLogExportRequest struct {
	// The name of the exported file. The name can be up to 1,024 characters in length. The default name suffix is log.csv.
	//
	// example:
	//
	// 2024-01-log.csv
	FileName *string `json:"file_name,omitempty" xml:"file_name,omitempty"`
	// The export language. Default value: zh-CN. Valid values:
	//
	// 	- zh-CN: Chinese
	//
	// 	- en_US: English
	//
	// example:
	//
	// zh_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// The sort order based on the operation time. If you leave this parameter empty, the value acted_at DESC is used. Valid values:
	//
	// 	- acted_at DESC: sorts the entries by operation time in descending order
	//
	// 	- acted_at ASC: sorts the entries by operation time in ascending order
	//
	// example:
	//
	// acted_at DESC
	OrderBy *string `json:"order_by,omitempty" xml:"order_by,omitempty"`
	// The fields used for query. You can specify one or more of the following fields:
	//
	// 	- drive_id (space ID, in the form of a string)
	//
	// 	- actor_id (operator ID, in the form of a string)
	//
	// 	- acted_at (operation time, in the yyyy-MM-ddTHH:mm:ssZ format in UTC, for example, 2006-01-02T00:00:00)
	//
	// 	- action_type (operation type, in the form of a string)
	//
	// example:
	//
	// acted_at > \\"2025-03-10T16:00:00\\" and acted_at < \\"2025-03-17T15:59:59\\"
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
}

func (s AuditLogExportRequest) String() string {
	return dara.Prettify(s)
}

func (s AuditLogExportRequest) GoString() string {
	return s.String()
}

func (s *AuditLogExportRequest) GetFileName() *string {
	return s.FileName
}

func (s *AuditLogExportRequest) GetLanguage() *string {
	return s.Language
}

func (s *AuditLogExportRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *AuditLogExportRequest) GetQuery() *string {
	return s.Query
}

func (s *AuditLogExportRequest) SetFileName(v string) *AuditLogExportRequest {
	s.FileName = &v
	return s
}

func (s *AuditLogExportRequest) SetLanguage(v string) *AuditLogExportRequest {
	s.Language = &v
	return s
}

func (s *AuditLogExportRequest) SetOrderBy(v string) *AuditLogExportRequest {
	s.OrderBy = &v
	return s
}

func (s *AuditLogExportRequest) SetQuery(v string) *AuditLogExportRequest {
	s.Query = &v
	return s
}

func (s *AuditLogExportRequest) Validate() error {
	return dara.Validate(s)
}
