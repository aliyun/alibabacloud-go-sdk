// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainOrderQueryV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TrainOrderQueryV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TrainOrderQueryV2Response
	GetStatusCode() *int32
	SetBody(v *TrainOrderQueryV2ResponseBody) *TrainOrderQueryV2Response
	GetBody() *TrainOrderQueryV2ResponseBody
}

type TrainOrderQueryV2Response struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TrainOrderQueryV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TrainOrderQueryV2Response) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderQueryV2Response) GoString() string {
	return s.String()
}

func (s *TrainOrderQueryV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TrainOrderQueryV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TrainOrderQueryV2Response) GetBody() *TrainOrderQueryV2ResponseBody {
	return s.Body
}

func (s *TrainOrderQueryV2Response) SetHeaders(v map[string]*string) *TrainOrderQueryV2Response {
	s.Headers = v
	return s
}

func (s *TrainOrderQueryV2Response) SetStatusCode(v int32) *TrainOrderQueryV2Response {
	s.StatusCode = &v
	return s
}

func (s *TrainOrderQueryV2Response) SetBody(v *TrainOrderQueryV2ResponseBody) *TrainOrderQueryV2Response {
	s.Body = v
	return s
}

func (s *TrainOrderQueryV2Response) Validate() error {
	return dara.Validate(s)
}
