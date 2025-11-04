// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCopyrightExtractJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *QueryCopyrightExtractJobResponseBodyData) *QueryCopyrightExtractJobResponseBody
	GetData() *QueryCopyrightExtractJobResponseBodyData
	SetMessage(v string) *QueryCopyrightExtractJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryCopyrightExtractJobResponseBody
	GetRequestId() *string
	SetStatusCode(v int64) *QueryCopyrightExtractJobResponseBody
	GetStatusCode() *int64
}

type QueryCopyrightExtractJobResponseBody struct {
	// The data returned.
	Data *QueryCopyrightExtractJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// *****ACB-44F2-5F2D-88D7-1283E70*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status code.
	//
	// example:
	//
	// 200
	StatusCode *int64 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s QueryCopyrightExtractJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCopyrightExtractJobResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCopyrightExtractJobResponseBody) GetData() *QueryCopyrightExtractJobResponseBodyData {
	return s.Data
}

func (s *QueryCopyrightExtractJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryCopyrightExtractJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCopyrightExtractJobResponseBody) GetStatusCode() *int64 {
	return s.StatusCode
}

func (s *QueryCopyrightExtractJobResponseBody) SetData(v *QueryCopyrightExtractJobResponseBodyData) *QueryCopyrightExtractJobResponseBody {
	s.Data = v
	return s
}

func (s *QueryCopyrightExtractJobResponseBody) SetMessage(v string) *QueryCopyrightExtractJobResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCopyrightExtractJobResponseBody) SetRequestId(v string) *QueryCopyrightExtractJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCopyrightExtractJobResponseBody) SetStatusCode(v int64) *QueryCopyrightExtractJobResponseBody {
	s.StatusCode = &v
	return s
}

func (s *QueryCopyrightExtractJobResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryCopyrightExtractJobResponseBodyData struct {
	// The copyright watermark information.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s QueryCopyrightExtractJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryCopyrightExtractJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryCopyrightExtractJobResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *QueryCopyrightExtractJobResponseBodyData) SetMessage(v string) *QueryCopyrightExtractJobResponseBodyData {
	s.Message = &v
	return s
}

func (s *QueryCopyrightExtractJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
