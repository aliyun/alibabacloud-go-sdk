// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFileTaskResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *FileTaskResultResponse
	GetErrCode() *string
	SetMessage(v string) *FileTaskResultResponse
	GetMessage() *string
	SetRstFile(v *FileIDInfo) *FileTaskResultResponse
	GetRstFile() *FileIDInfo
	SetSrcFile(v *FileIDInfo) *FileTaskResultResponse
	GetSrcFile() *FileIDInfo
}

type FileTaskResultResponse struct {
	ErrCode *string     `json:"err_code,omitempty" xml:"err_code,omitempty"`
	Message *string     `json:"message,omitempty" xml:"message,omitempty"`
	RstFile *FileIDInfo `json:"rst_file,omitempty" xml:"rst_file,omitempty"`
	SrcFile *FileIDInfo `json:"src_file,omitempty" xml:"src_file,omitempty"`
}

func (s FileTaskResultResponse) String() string {
	return dara.Prettify(s)
}

func (s FileTaskResultResponse) GoString() string {
	return s.String()
}

func (s *FileTaskResultResponse) GetErrCode() *string {
	return s.ErrCode
}

func (s *FileTaskResultResponse) GetMessage() *string {
	return s.Message
}

func (s *FileTaskResultResponse) GetRstFile() *FileIDInfo {
	return s.RstFile
}

func (s *FileTaskResultResponse) GetSrcFile() *FileIDInfo {
	return s.SrcFile
}

func (s *FileTaskResultResponse) SetErrCode(v string) *FileTaskResultResponse {
	s.ErrCode = &v
	return s
}

func (s *FileTaskResultResponse) SetMessage(v string) *FileTaskResultResponse {
	s.Message = &v
	return s
}

func (s *FileTaskResultResponse) SetRstFile(v *FileIDInfo) *FileTaskResultResponse {
	s.RstFile = v
	return s
}

func (s *FileTaskResultResponse) SetSrcFile(v *FileIDInfo) *FileTaskResultResponse {
	s.SrcFile = v
	return s
}

func (s *FileTaskResultResponse) Validate() error {
	if s.RstFile != nil {
		if err := s.RstFile.Validate(); err != nil {
			return err
		}
	}
	if s.SrcFile != nil {
		if err := s.SrcFile.Validate(); err != nil {
			return err
		}
	}
	return nil
}
