// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImportJobResult(v *GetJobResultResp) *GetJobResultResponseBody
	GetImportJobResult() *GetJobResultResp
}

type GetJobResultResponseBody struct {
	// The details for obtaining the retries of the migration task.
	ImportJobResult *GetJobResultResp `json:"ImportJobResult,omitempty" xml:"ImportJobResult,omitempty"`
}

func (s GetJobResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetJobResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobResultResponseBody) GetImportJobResult() *GetJobResultResp {
	return s.ImportJobResult
}

func (s *GetJobResultResponseBody) SetImportJobResult(v *GetJobResultResp) *GetJobResultResponseBody {
	s.ImportJobResult = v
	return s
}

func (s *GetJobResultResponseBody) Validate() error {
	if s.ImportJobResult != nil {
		if err := s.ImportJobResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}
