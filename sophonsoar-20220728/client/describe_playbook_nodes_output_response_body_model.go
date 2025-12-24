// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlaybookNodesOutputResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPlaybookNodesOutput(v *DescribePlaybookNodesOutputResponseBodyPlaybookNodesOutput) *DescribePlaybookNodesOutputResponseBody
	GetPlaybookNodesOutput() *DescribePlaybookNodesOutputResponseBodyPlaybookNodesOutput
	SetRequestId(v string) *DescribePlaybookNodesOutputResponseBody
	GetRequestId() *string
}

type DescribePlaybookNodesOutputResponseBody struct {
	// The output data of the component node.
	PlaybookNodesOutput *DescribePlaybookNodesOutputResponseBodyPlaybookNodesOutput `json:"PlaybookNodesOutput,omitempty" xml:"PlaybookNodesOutput,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A491170C-FE1F-520E-83D4-72ED205B72ED
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePlaybookNodesOutputResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePlaybookNodesOutputResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePlaybookNodesOutputResponseBody) GetPlaybookNodesOutput() *DescribePlaybookNodesOutputResponseBodyPlaybookNodesOutput {
	return s.PlaybookNodesOutput
}

func (s *DescribePlaybookNodesOutputResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePlaybookNodesOutputResponseBody) SetPlaybookNodesOutput(v *DescribePlaybookNodesOutputResponseBodyPlaybookNodesOutput) *DescribePlaybookNodesOutputResponseBody {
	s.PlaybookNodesOutput = v
	return s
}

func (s *DescribePlaybookNodesOutputResponseBody) SetRequestId(v string) *DescribePlaybookNodesOutputResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePlaybookNodesOutputResponseBody) Validate() error {
	if s.PlaybookNodesOutput != nil {
		if err := s.PlaybookNodesOutput.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePlaybookNodesOutputResponseBodyPlaybookNodesOutput struct {
	// The name of the component node.
	//
	// example:
	//
	// DataFormat_1
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The historical output data of the component node. The value is in the JSON string format. If no data is found, the parameter is left empty.
	//
	// example:
	//
	// {
	//
	//     "datalist": [
	//
	//         {
	//
	//             "score": "10",
	//
	//             "ip": "1.1.1.1"
	//
	//         }
	//
	//     ],
	//
	//     "total_data_successful": 1,
	//
	//     "filter_total_data": 1,
	//
	//     "total_data": 1,
	//
	//     "total_exe_successful": 1,
	//
	//     "total_exe": 1,
	//
	//     "total_data_with_dup": 1,
	//
	//     "filter_total_data_successful": 1,
	//
	//     "status": true
	//
	// }
	NodeOutput *string `json:"NodeOutput,omitempty" xml:"NodeOutput,omitempty"`
}

func (s DescribePlaybookNodesOutputResponseBodyPlaybookNodesOutput) String() string {
	return dara.Prettify(s)
}

func (s DescribePlaybookNodesOutputResponseBodyPlaybookNodesOutput) GoString() string {
	return s.String()
}

func (s *DescribePlaybookNodesOutputResponseBodyPlaybookNodesOutput) GetNodeName() *string {
	return s.NodeName
}

func (s *DescribePlaybookNodesOutputResponseBodyPlaybookNodesOutput) GetNodeOutput() *string {
	return s.NodeOutput
}

func (s *DescribePlaybookNodesOutputResponseBodyPlaybookNodesOutput) SetNodeName(v string) *DescribePlaybookNodesOutputResponseBodyPlaybookNodesOutput {
	s.NodeName = &v
	return s
}

func (s *DescribePlaybookNodesOutputResponseBodyPlaybookNodesOutput) SetNodeOutput(v string) *DescribePlaybookNodesOutputResponseBodyPlaybookNodesOutput {
	s.NodeOutput = &v
	return s
}

func (s *DescribePlaybookNodesOutputResponseBodyPlaybookNodesOutput) Validate() error {
	return dara.Validate(s)
}
