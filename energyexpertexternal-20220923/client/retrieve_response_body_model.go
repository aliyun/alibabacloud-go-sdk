// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetrieveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *RetrieveResponseBodyData) *RetrieveResponseBody
	GetData() *RetrieveResponseBodyData
	SetRequestId(v string) *RetrieveResponseBody
	GetRequestId() *string
}

type RetrieveResponseBody struct {
	Data *RetrieveResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RetrieveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RetrieveResponseBody) GoString() string {
	return s.String()
}

func (s *RetrieveResponseBody) GetData() *RetrieveResponseBodyData {
	return s.Data
}

func (s *RetrieveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RetrieveResponseBody) SetData(v *RetrieveResponseBodyData) *RetrieveResponseBody {
	s.Data = v
	return s
}

func (s *RetrieveResponseBody) SetRequestId(v string) *RetrieveResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetrieveResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RetrieveResponseBodyData struct {
	// example:
	//
	// 76cd1691-daf6-4113-8444-4998a25a5323
	JobId   *string      `json:"jobId,omitempty" xml:"jobId,omitempty"`
	Results []*ChunkItem `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
}

func (s RetrieveResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RetrieveResponseBodyData) GoString() string {
	return s.String()
}

func (s *RetrieveResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *RetrieveResponseBodyData) GetResults() []*ChunkItem {
	return s.Results
}

func (s *RetrieveResponseBodyData) SetJobId(v string) *RetrieveResponseBodyData {
	s.JobId = &v
	return s
}

func (s *RetrieveResponseBodyData) SetResults(v []*ChunkItem) *RetrieveResponseBodyData {
	s.Results = v
	return s
}

func (s *RetrieveResponseBodyData) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
