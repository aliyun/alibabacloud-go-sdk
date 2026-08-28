// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPatrolReportDetailHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetPatrolReportDetailHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *GetPatrolReportDetailHeaders
	GetWorkspace() *string
}

type GetPatrolReportDetailHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetPatrolReportDetailHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetPatrolReportDetailHeaders) GoString() string {
	return s.String()
}

func (s *GetPatrolReportDetailHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetPatrolReportDetailHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetPatrolReportDetailHeaders) SetCommonHeaders(v map[string]*string) *GetPatrolReportDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetPatrolReportDetailHeaders) SetWorkspace(v string) *GetPatrolReportDetailHeaders {
	s.Workspace = &v
	return s
}

func (s *GetPatrolReportDetailHeaders) Validate() error {
	return dara.Validate(s)
}
