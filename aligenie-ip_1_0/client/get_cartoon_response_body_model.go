// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCartoonResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *GetCartoonResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCartoonResponseBody
	GetRequestId() *string
	SetResult(v *GetCartoonResponseBodyResult) *GetCartoonResponseBody
	GetResult() *GetCartoonResponseBodyResult
	SetStatusCode(v int32) *GetCartoonResponseBody
	GetStatusCode() *int32
}

type GetCartoonResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7*726E
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetCartoonResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s GetCartoonResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCartoonResponseBody) GoString() string {
	return s.String()
}

func (s *GetCartoonResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCartoonResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCartoonResponseBody) GetResult() *GetCartoonResponseBodyResult {
	return s.Result
}

func (s *GetCartoonResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCartoonResponseBody) SetMessage(v string) *GetCartoonResponseBody {
	s.Message = &v
	return s
}

func (s *GetCartoonResponseBody) SetRequestId(v string) *GetCartoonResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCartoonResponseBody) SetResult(v *GetCartoonResponseBodyResult) *GetCartoonResponseBody {
	s.Result = v
	return s
}

func (s *GetCartoonResponseBody) SetStatusCode(v int32) *GetCartoonResponseBody {
	s.StatusCode = &v
	return s
}

func (s *GetCartoonResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCartoonResponseBodyResult struct {
	// example:
	//
	// https://ai***.mp4
	StartVideoMd5 *string `json:"StartVideoMd5,omitempty" xml:"StartVideoMd5,omitempty"`
	// example:
	//
	// 40c8***97
	StartVideoUrl *string `json:"StartVideoUrl,omitempty" xml:"StartVideoUrl,omitempty"`
}

func (s GetCartoonResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetCartoonResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetCartoonResponseBodyResult) GetStartVideoMd5() *string {
	return s.StartVideoMd5
}

func (s *GetCartoonResponseBodyResult) GetStartVideoUrl() *string {
	return s.StartVideoUrl
}

func (s *GetCartoonResponseBodyResult) SetStartVideoMd5(v string) *GetCartoonResponseBodyResult {
	s.StartVideoMd5 = &v
	return s
}

func (s *GetCartoonResponseBodyResult) SetStartVideoUrl(v string) *GetCartoonResponseBodyResult {
	s.StartVideoUrl = &v
	return s
}

func (s *GetCartoonResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
