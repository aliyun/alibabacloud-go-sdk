// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAISearchStreamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *AISearchStreamItem) *AISearchStreamResponseBody
	GetData() *AISearchStreamItem
	SetRequestId(v string) *AISearchStreamResponseBody
	GetRequestId() *string
}

type AISearchStreamResponseBody struct {
	Data *AISearchStreamItem `json:"data,omitempty" xml:"data,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 3F21E33A-42F8-50BC-89DE-DC0B96B967DD
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AISearchStreamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AISearchStreamResponseBody) GoString() string {
	return s.String()
}

func (s *AISearchStreamResponseBody) GetData() *AISearchStreamItem {
	return s.Data
}

func (s *AISearchStreamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AISearchStreamResponseBody) SetData(v *AISearchStreamItem) *AISearchStreamResponseBody {
	s.Data = v
	return s
}

func (s *AISearchStreamResponseBody) SetRequestId(v string) *AISearchStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *AISearchStreamResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
