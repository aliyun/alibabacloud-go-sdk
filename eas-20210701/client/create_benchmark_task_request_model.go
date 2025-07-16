// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBenchmarkTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *CreateBenchmarkTaskRequest
	GetBody() *string
}

type CreateBenchmarkTaskRequest struct {
	// The request body. The body includes the parameters that are set to create a stress testing task.
	//
	// example:
	//
	// {
	//
	//     "base": {
	//
	//         "duration": 600
	//
	//     },
	//
	//     "service": {
	//
	//         "serviceName": "test_service",
	//
	//         "requestToken": "test_token"
	//
	//     },
	//
	//     "data": {
	//
	//         "path": "https://larec-benchmark-cd.oss-cn-chengdu.aliyuncs.com/youbei/sv_dbmtl/data/youbei.warmup.tf.bin",
	//
	//         "dataType": "binary"
	//
	//     },
	//
	//     "optional": {
	//
	//        "maxRt": 100
	//
	//     }
	//
	// }
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBenchmarkTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBenchmarkTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateBenchmarkTaskRequest) GetBody() *string {
	return s.Body
}

func (s *CreateBenchmarkTaskRequest) SetBody(v string) *CreateBenchmarkTaskRequest {
	s.Body = &v
	return s
}

func (s *CreateBenchmarkTaskRequest) Validate() error {
	return dara.Validate(s)
}
