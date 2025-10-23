// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetValidateFileStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCatchAllNum(v string) *GetValidateFileStatusResponseBody
	GetCatchAllNum() *string
	SetCompleteTime(v string) *GetValidateFileStatusResponseBody
	GetCompleteTime() *string
	SetDoNotMailNum(v string) *GetValidateFileStatusResponseBody
	GetDoNotMailNum() *string
	SetFileName(v string) *GetValidateFileStatusResponseBody
	GetFileName() *string
	SetInvalidNum(v string) *GetValidateFileStatusResponseBody
	GetInvalidNum() *string
	SetPercentage(v string) *GetValidateFileStatusResponseBody
	GetPercentage() *string
	SetProcessedNum(v string) *GetValidateFileStatusResponseBody
	GetProcessedNum() *string
	SetRequestId(v string) *GetValidateFileStatusResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetValidateFileStatusResponseBody
	GetStatus() *string
	SetTotalNum(v string) *GetValidateFileStatusResponseBody
	GetTotalNum() *string
	SetUnknownNum(v string) *GetValidateFileStatusResponseBody
	GetUnknownNum() *string
	SetUploadTime(v string) *GetValidateFileStatusResponseBody
	GetUploadTime() *string
	SetValidNum(v string) *GetValidateFileStatusResponseBody
	GetValidNum() *string
}

type GetValidateFileStatusResponseBody struct {
	// example:
	//
	// 2
	CatchAllNum *string `json:"CatchAllNum,omitempty" xml:"CatchAllNum,omitempty"`
	// example:
	//
	// 2000-01-01T00:00:00Z
	CompleteTime *string `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// example:
	//
	// 1
	DoNotMailNum *string `json:"DoNotMailNum,omitempty" xml:"DoNotMailNum,omitempty"`
	// example:
	//
	// file.txt
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// 2
	InvalidNum *string `json:"InvalidNum,omitempty" xml:"InvalidNum,omitempty"`
	// example:
	//
	// 100%
	Percentage *string `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	// example:
	//
	// 10
	ProcessedNum *string `json:"ProcessedNum,omitempty" xml:"ProcessedNum,omitempty"`
	// example:
	//
	// yyyy-yyyy-yyyy-yyyy
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// completed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 10
	TotalNum *string `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// example:
	//
	// 1
	UnknownNum *string `json:"UnknownNum,omitempty" xml:"UnknownNum,omitempty"`
	// example:
	//
	// 2000-01-01T00:00:00Z
	UploadTime *string `json:"UploadTime,omitempty" xml:"UploadTime,omitempty"`
	// example:
	//
	// 4
	ValidNum *string `json:"ValidNum,omitempty" xml:"ValidNum,omitempty"`
}

func (s GetValidateFileStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetValidateFileStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetValidateFileStatusResponseBody) GetCatchAllNum() *string {
	return s.CatchAllNum
}

func (s *GetValidateFileStatusResponseBody) GetCompleteTime() *string {
	return s.CompleteTime
}

func (s *GetValidateFileStatusResponseBody) GetDoNotMailNum() *string {
	return s.DoNotMailNum
}

func (s *GetValidateFileStatusResponseBody) GetFileName() *string {
	return s.FileName
}

func (s *GetValidateFileStatusResponseBody) GetInvalidNum() *string {
	return s.InvalidNum
}

func (s *GetValidateFileStatusResponseBody) GetPercentage() *string {
	return s.Percentage
}

func (s *GetValidateFileStatusResponseBody) GetProcessedNum() *string {
	return s.ProcessedNum
}

func (s *GetValidateFileStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetValidateFileStatusResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetValidateFileStatusResponseBody) GetTotalNum() *string {
	return s.TotalNum
}

func (s *GetValidateFileStatusResponseBody) GetUnknownNum() *string {
	return s.UnknownNum
}

func (s *GetValidateFileStatusResponseBody) GetUploadTime() *string {
	return s.UploadTime
}

func (s *GetValidateFileStatusResponseBody) GetValidNum() *string {
	return s.ValidNum
}

func (s *GetValidateFileStatusResponseBody) SetCatchAllNum(v string) *GetValidateFileStatusResponseBody {
	s.CatchAllNum = &v
	return s
}

func (s *GetValidateFileStatusResponseBody) SetCompleteTime(v string) *GetValidateFileStatusResponseBody {
	s.CompleteTime = &v
	return s
}

func (s *GetValidateFileStatusResponseBody) SetDoNotMailNum(v string) *GetValidateFileStatusResponseBody {
	s.DoNotMailNum = &v
	return s
}

func (s *GetValidateFileStatusResponseBody) SetFileName(v string) *GetValidateFileStatusResponseBody {
	s.FileName = &v
	return s
}

func (s *GetValidateFileStatusResponseBody) SetInvalidNum(v string) *GetValidateFileStatusResponseBody {
	s.InvalidNum = &v
	return s
}

func (s *GetValidateFileStatusResponseBody) SetPercentage(v string) *GetValidateFileStatusResponseBody {
	s.Percentage = &v
	return s
}

func (s *GetValidateFileStatusResponseBody) SetProcessedNum(v string) *GetValidateFileStatusResponseBody {
	s.ProcessedNum = &v
	return s
}

func (s *GetValidateFileStatusResponseBody) SetRequestId(v string) *GetValidateFileStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetValidateFileStatusResponseBody) SetStatus(v string) *GetValidateFileStatusResponseBody {
	s.Status = &v
	return s
}

func (s *GetValidateFileStatusResponseBody) SetTotalNum(v string) *GetValidateFileStatusResponseBody {
	s.TotalNum = &v
	return s
}

func (s *GetValidateFileStatusResponseBody) SetUnknownNum(v string) *GetValidateFileStatusResponseBody {
	s.UnknownNum = &v
	return s
}

func (s *GetValidateFileStatusResponseBody) SetUploadTime(v string) *GetValidateFileStatusResponseBody {
	s.UploadTime = &v
	return s
}

func (s *GetValidateFileStatusResponseBody) SetValidNum(v string) *GetValidateFileStatusResponseBody {
	s.ValidNum = &v
	return s
}

func (s *GetValidateFileStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
