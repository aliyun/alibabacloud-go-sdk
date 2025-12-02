// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobNameListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*string) *GetJobNameListResponseBody
	GetData() []*string
	SetRequestId(v string) *GetJobNameListResponseBody
	GetRequestId() *string
}

type GetJobNameListResponseBody struct {
	// Returned data.
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// ID assigned by the backend, used to uniquely identify a request. Can be used for troubleshooting.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetJobNameListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetJobNameListResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobNameListResponseBody) GetData() []*string {
	return s.Data
}

func (s *GetJobNameListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetJobNameListResponseBody) SetData(v []*string) *GetJobNameListResponseBody {
	s.Data = v
	return s
}

func (s *GetJobNameListResponseBody) SetRequestId(v string) *GetJobNameListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobNameListResponseBody) Validate() error {
	return dara.Validate(s)
}
