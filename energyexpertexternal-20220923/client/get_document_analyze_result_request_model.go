// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentAnalyzeResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetDocumentAnalyzeResultRequest
	GetJobId() *string
}

type GetDocumentAnalyzeResultRequest struct {
	// Job ID, specifying which document\\"s parsing result to query. This is a return parameter from the \\"Submit Document Parsing Job\\" interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// bfce2248-1546-4298-8bcf-70ac26e69646
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
}

func (s GetDocumentAnalyzeResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentAnalyzeResultRequest) GoString() string {
	return s.String()
}

func (s *GetDocumentAnalyzeResultRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetDocumentAnalyzeResultRequest) SetJobId(v string) *GetDocumentAnalyzeResultRequest {
	s.JobId = &v
	return s
}

func (s *GetDocumentAnalyzeResultRequest) Validate() error {
	return dara.Validate(s)
}
