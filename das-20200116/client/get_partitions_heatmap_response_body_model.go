// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPartitionsHeatmapResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetPartitionsHeatmapResponseBody
	GetCode() *string
	SetData(v string) *GetPartitionsHeatmapResponseBody
	GetData() *string
	SetMessage(v string) *GetPartitionsHeatmapResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPartitionsHeatmapResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetPartitionsHeatmapResponseBody
	GetSuccess() *string
}

type GetPartitionsHeatmapResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The hot data of the PolarDB-X 2.0 instance. The data is returned in JSON format.
	//
	// example:
	//
	// {
	//
	//     "boundAxis": [
	//
	//         {
	//
	//             "bound": "A,B,C,D",
	//
	//             "labels": [
	//
	//                 "L1",
	//
	//                 "L2",
	//
	//                 "L3",
	//
	//                 "L4"
	//
	//             ],
	//
	//             "rows": 3171
	//
	//         },
	//
	//         {
	//
	//             "bound": "A,B,C,D",
	//
	//             "labels": [
	//
	//                 "L1",
	//
	//                 "L2",
	//
	//                 "L3",
	//
	//                 "L4"
	//
	//             ],
	//
	//             "rows": 277128
	//
	//         }
	//
	//     ],
	//
	//     "dataMap": {
	//
	//         "READ_WRITTEN_ROWS": [
	//
	//             [
	//
	//                 0,
	//
	//                 0,
	//
	//                 0
	//
	//             ],
	//
	//             [
	//
	//                 0,
	//
	//                 0,
	//
	//                 0
	//
	//             ]
	//
	//         ]
	//
	//     },
	//
	//     "timeAxis": [
	//
	//         1671701056070,
	//
	//         1671701116551,
	//
	//         1671701177020
	//
	//     ]
	//
	// }
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// >  If the request was successful, **Successful*	- is returned. If the request failed, an error message that contains information such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D00DB161-FEF6-5428-B37A-8D29A4C2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPartitionsHeatmapResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPartitionsHeatmapResponseBody) GoString() string {
	return s.String()
}

func (s *GetPartitionsHeatmapResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPartitionsHeatmapResponseBody) GetData() *string {
	return s.Data
}

func (s *GetPartitionsHeatmapResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPartitionsHeatmapResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPartitionsHeatmapResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetPartitionsHeatmapResponseBody) SetCode(v string) *GetPartitionsHeatmapResponseBody {
	s.Code = &v
	return s
}

func (s *GetPartitionsHeatmapResponseBody) SetData(v string) *GetPartitionsHeatmapResponseBody {
	s.Data = &v
	return s
}

func (s *GetPartitionsHeatmapResponseBody) SetMessage(v string) *GetPartitionsHeatmapResponseBody {
	s.Message = &v
	return s
}

func (s *GetPartitionsHeatmapResponseBody) SetRequestId(v string) *GetPartitionsHeatmapResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPartitionsHeatmapResponseBody) SetSuccess(v string) *GetPartitionsHeatmapResponseBody {
	s.Success = &v
	return s
}

func (s *GetPartitionsHeatmapResponseBody) Validate() error {
	return dara.Validate(s)
}
