// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkitemFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetWorkitemFileResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GetWorkitemFileResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *GetWorkitemFileResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetWorkitemFileResponseBody
	GetSuccess() *string
	SetWorkitemFile(v *GetWorkitemFileResponseBodyWorkitemFile) *GetWorkitemFileResponseBody
	GetWorkitemFile() *GetWorkitemFileResponseBodyWorkitemFile
}

type GetWorkitemFileResponseBody struct {
	// example:
	//
	// success
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// error
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// Id of the request
	//
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success      *string                                  `json:"success,omitempty" xml:"success,omitempty"`
	WorkitemFile *GetWorkitemFileResponseBodyWorkitemFile `json:"workitemFile,omitempty" xml:"workitemFile,omitempty" type:"Struct"`
}

func (s GetWorkitemFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkitemFileResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkitemFileResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetWorkitemFileResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetWorkitemFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkitemFileResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetWorkitemFileResponseBody) GetWorkitemFile() *GetWorkitemFileResponseBodyWorkitemFile {
	return s.WorkitemFile
}

func (s *GetWorkitemFileResponseBody) SetErrorCode(v string) *GetWorkitemFileResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetWorkitemFileResponseBody) SetErrorMsg(v string) *GetWorkitemFileResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetWorkitemFileResponseBody) SetRequestId(v string) *GetWorkitemFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkitemFileResponseBody) SetSuccess(v string) *GetWorkitemFileResponseBody {
	s.Success = &v
	return s
}

func (s *GetWorkitemFileResponseBody) SetWorkitemFile(v *GetWorkitemFileResponseBodyWorkitemFile) *GetWorkitemFileResponseBody {
	s.WorkitemFile = v
	return s
}

func (s *GetWorkitemFileResponseBody) Validate() error {
	if s.WorkitemFile != nil {
		if err := s.WorkitemFile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkitemFileResponseBodyWorkitemFile struct {
	// example:
	//
	// sddrdfdf123df
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// ddc.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 10001
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// pdf
	Suffix *string `json:"suffix,omitempty" xml:"suffix,omitempty"`
	// example:
	//
	// http://tmaestro-oss.oss-cn-hongkong.aliyuncs.com/thread_1682129288279.log
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetWorkitemFileResponseBodyWorkitemFile) String() string {
	return dara.Prettify(s)
}

func (s GetWorkitemFileResponseBodyWorkitemFile) GoString() string {
	return s.String()
}

func (s *GetWorkitemFileResponseBodyWorkitemFile) GetId() *string {
	return s.Id
}

func (s *GetWorkitemFileResponseBodyWorkitemFile) GetName() *string {
	return s.Name
}

func (s *GetWorkitemFileResponseBodyWorkitemFile) GetSize() *int32 {
	return s.Size
}

func (s *GetWorkitemFileResponseBodyWorkitemFile) GetSuffix() *string {
	return s.Suffix
}

func (s *GetWorkitemFileResponseBodyWorkitemFile) GetUrl() *string {
	return s.Url
}

func (s *GetWorkitemFileResponseBodyWorkitemFile) SetId(v string) *GetWorkitemFileResponseBodyWorkitemFile {
	s.Id = &v
	return s
}

func (s *GetWorkitemFileResponseBodyWorkitemFile) SetName(v string) *GetWorkitemFileResponseBodyWorkitemFile {
	s.Name = &v
	return s
}

func (s *GetWorkitemFileResponseBodyWorkitemFile) SetSize(v int32) *GetWorkitemFileResponseBodyWorkitemFile {
	s.Size = &v
	return s
}

func (s *GetWorkitemFileResponseBodyWorkitemFile) SetSuffix(v string) *GetWorkitemFileResponseBodyWorkitemFile {
	s.Suffix = &v
	return s
}

func (s *GetWorkitemFileResponseBodyWorkitemFile) SetUrl(v string) *GetWorkitemFileResponseBodyWorkitemFile {
	s.Url = &v
	return s
}

func (s *GetWorkitemFileResponseBodyWorkitemFile) Validate() error {
	return dara.Validate(s)
}
