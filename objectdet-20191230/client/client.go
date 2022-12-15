// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	openplatform "github.com/alibabacloud-go/openplatform-20191219/v2/client"
	fileform "github.com/alibabacloud-go/tea-fileform/service"
	oss "github.com/alibabacloud-go/tea-oss-sdk/client"
	ossutil "github.com/alibabacloud-go/tea-oss-utils/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"io"
)

type DetectObjectElement struct {
	Height *int64   `json:"Height,omitempty" xml:"Height,omitempty"`
	Score  *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	Type   *string  `json:"Type,omitempty" xml:"Type,omitempty"`
	Width  *int64   `json:"Width,omitempty" xml:"Width,omitempty"`
	X      *int64   `json:"X,omitempty" xml:"X,omitempty"`
	Y      *int64   `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DetectObjectElement) String() string {
	return tea.Prettify(s)
}

func (s DetectObjectElement) GoString() string {
	return s.String()
}

func (s *DetectObjectElement) SetHeight(v int64) *DetectObjectElement {
	s.Height = &v
	return s
}

func (s *DetectObjectElement) SetScore(v float32) *DetectObjectElement {
	s.Score = &v
	return s
}

func (s *DetectObjectElement) SetType(v string) *DetectObjectElement {
	s.Type = &v
	return s
}

func (s *DetectObjectElement) SetWidth(v int64) *DetectObjectElement {
	s.Width = &v
	return s
}

func (s *DetectObjectElement) SetX(v int64) *DetectObjectElement {
	s.X = &v
	return s
}

func (s *DetectObjectElement) SetY(v int64) *DetectObjectElement {
	s.Y = &v
	return s
}

type DetectObjectFrame struct {
	Elements []*DetectObjectElement `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	Time     *int64                 `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DetectObjectFrame) String() string {
	return tea.Prettify(s)
}

func (s DetectObjectFrame) GoString() string {
	return s.String()
}

func (s *DetectObjectFrame) SetElements(v []*DetectObjectElement) *DetectObjectFrame {
	s.Elements = v
	return s
}

func (s *DetectObjectFrame) SetTime(v int64) *DetectObjectFrame {
	s.Time = &v
	return s
}

type ClassifyVehicleInsuranceRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s ClassifyVehicleInsuranceRequest) String() string {
	return tea.Prettify(s)
}

func (s ClassifyVehicleInsuranceRequest) GoString() string {
	return s.String()
}

func (s *ClassifyVehicleInsuranceRequest) SetImageURL(v string) *ClassifyVehicleInsuranceRequest {
	s.ImageURL = &v
	return s
}

type ClassifyVehicleInsuranceAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s ClassifyVehicleInsuranceAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ClassifyVehicleInsuranceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ClassifyVehicleInsuranceAdvanceRequest) SetImageURLObject(v io.Reader) *ClassifyVehicleInsuranceAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type ClassifyVehicleInsuranceResponseBody struct {
	Data      *ClassifyVehicleInsuranceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClassifyVehicleInsuranceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ClassifyVehicleInsuranceResponseBody) GoString() string {
	return s.String()
}

func (s *ClassifyVehicleInsuranceResponseBody) SetData(v *ClassifyVehicleInsuranceResponseBodyData) *ClassifyVehicleInsuranceResponseBody {
	s.Data = v
	return s
}

func (s *ClassifyVehicleInsuranceResponseBody) SetRequestId(v string) *ClassifyVehicleInsuranceResponseBody {
	s.RequestId = &v
	return s
}

type ClassifyVehicleInsuranceResponseBodyData struct {
	Labels    []*ClassifyVehicleInsuranceResponseBodyDataLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Threshold *float32                                          `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s ClassifyVehicleInsuranceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ClassifyVehicleInsuranceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ClassifyVehicleInsuranceResponseBodyData) SetLabels(v []*ClassifyVehicleInsuranceResponseBodyDataLabels) *ClassifyVehicleInsuranceResponseBodyData {
	s.Labels = v
	return s
}

func (s *ClassifyVehicleInsuranceResponseBodyData) SetThreshold(v float32) *ClassifyVehicleInsuranceResponseBodyData {
	s.Threshold = &v
	return s
}

type ClassifyVehicleInsuranceResponseBodyDataLabels struct {
	Name  *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s ClassifyVehicleInsuranceResponseBodyDataLabels) String() string {
	return tea.Prettify(s)
}

func (s ClassifyVehicleInsuranceResponseBodyDataLabels) GoString() string {
	return s.String()
}

func (s *ClassifyVehicleInsuranceResponseBodyDataLabels) SetName(v string) *ClassifyVehicleInsuranceResponseBodyDataLabels {
	s.Name = &v
	return s
}

func (s *ClassifyVehicleInsuranceResponseBodyDataLabels) SetScore(v float32) *ClassifyVehicleInsuranceResponseBodyDataLabels {
	s.Score = &v
	return s
}

type ClassifyVehicleInsuranceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ClassifyVehicleInsuranceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ClassifyVehicleInsuranceResponse) String() string {
	return tea.Prettify(s)
}

func (s ClassifyVehicleInsuranceResponse) GoString() string {
	return s.String()
}

func (s *ClassifyVehicleInsuranceResponse) SetHeaders(v map[string]*string) *ClassifyVehicleInsuranceResponse {
	s.Headers = v
	return s
}

func (s *ClassifyVehicleInsuranceResponse) SetStatusCode(v int32) *ClassifyVehicleInsuranceResponse {
	s.StatusCode = &v
	return s
}

func (s *ClassifyVehicleInsuranceResponse) SetBody(v *ClassifyVehicleInsuranceResponseBody) *ClassifyVehicleInsuranceResponse {
	s.Body = v
	return s
}

type DetectIPCObjectRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectIPCObjectRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectIPCObjectRequest) GoString() string {
	return s.String()
}

func (s *DetectIPCObjectRequest) SetImageURL(v string) *DetectIPCObjectRequest {
	s.ImageURL = &v
	return s
}

type DetectIPCObjectAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectIPCObjectAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectIPCObjectAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectIPCObjectAdvanceRequest) SetImageURLObject(v io.Reader) *DetectIPCObjectAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type DetectIPCObjectResponseBody struct {
	Data      *DetectIPCObjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectIPCObjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectIPCObjectResponseBody) GoString() string {
	return s.String()
}

func (s *DetectIPCObjectResponseBody) SetData(v *DetectIPCObjectResponseBodyData) *DetectIPCObjectResponseBody {
	s.Data = v
	return s
}

func (s *DetectIPCObjectResponseBody) SetRequestId(v string) *DetectIPCObjectResponseBody {
	s.RequestId = &v
	return s
}

type DetectIPCObjectResponseBodyData struct {
	Elements []*DetectIPCObjectResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	Height   *int64                                     `json:"Height,omitempty" xml:"Height,omitempty"`
	Width    *int64                                     `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s DetectIPCObjectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectIPCObjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectIPCObjectResponseBodyData) SetElements(v []*DetectIPCObjectResponseBodyDataElements) *DetectIPCObjectResponseBodyData {
	s.Elements = v
	return s
}

func (s *DetectIPCObjectResponseBodyData) SetHeight(v int64) *DetectIPCObjectResponseBodyData {
	s.Height = &v
	return s
}

func (s *DetectIPCObjectResponseBodyData) SetWidth(v int64) *DetectIPCObjectResponseBodyData {
	s.Width = &v
	return s
}

type DetectIPCObjectResponseBodyDataElements struct {
	Box        []*int64 `json:"Box,omitempty" xml:"Box,omitempty" type:"Repeated"`
	Score      *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	TargetRate *float32 `json:"TargetRate,omitempty" xml:"TargetRate,omitempty"`
	Type       *string  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DetectIPCObjectResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s DetectIPCObjectResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *DetectIPCObjectResponseBodyDataElements) SetBox(v []*int64) *DetectIPCObjectResponseBodyDataElements {
	s.Box = v
	return s
}

func (s *DetectIPCObjectResponseBodyDataElements) SetScore(v float32) *DetectIPCObjectResponseBodyDataElements {
	s.Score = &v
	return s
}

func (s *DetectIPCObjectResponseBodyDataElements) SetTargetRate(v float32) *DetectIPCObjectResponseBodyDataElements {
	s.TargetRate = &v
	return s
}

func (s *DetectIPCObjectResponseBodyDataElements) SetType(v string) *DetectIPCObjectResponseBodyDataElements {
	s.Type = &v
	return s
}

type DetectIPCObjectResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectIPCObjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectIPCObjectResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectIPCObjectResponse) GoString() string {
	return s.String()
}

func (s *DetectIPCObjectResponse) SetHeaders(v map[string]*string) *DetectIPCObjectResponse {
	s.Headers = v
	return s
}

func (s *DetectIPCObjectResponse) SetStatusCode(v int32) *DetectIPCObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectIPCObjectResponse) SetBody(v *DetectIPCObjectResponseBody) *DetectIPCObjectResponse {
	s.Body = v
	return s
}

type DetectKitchenAnimalsRequest struct {
	ImageURLA *string `json:"ImageURLA,omitempty" xml:"ImageURLA,omitempty"`
	ImageURLB *string `json:"ImageURLB,omitempty" xml:"ImageURLB,omitempty"`
}

func (s DetectKitchenAnimalsRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectKitchenAnimalsRequest) GoString() string {
	return s.String()
}

func (s *DetectKitchenAnimalsRequest) SetImageURLA(v string) *DetectKitchenAnimalsRequest {
	s.ImageURLA = &v
	return s
}

func (s *DetectKitchenAnimalsRequest) SetImageURLB(v string) *DetectKitchenAnimalsRequest {
	s.ImageURLB = &v
	return s
}

type DetectKitchenAnimalsAdvanceRequest struct {
	ImageURLAObject io.Reader `json:"ImageURLA,omitempty" xml:"ImageURLA,omitempty"`
	ImageURLBObject io.Reader `json:"ImageURLB,omitempty" xml:"ImageURLB,omitempty"`
}

func (s DetectKitchenAnimalsAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectKitchenAnimalsAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectKitchenAnimalsAdvanceRequest) SetImageURLAObject(v io.Reader) *DetectKitchenAnimalsAdvanceRequest {
	s.ImageURLAObject = v
	return s
}

func (s *DetectKitchenAnimalsAdvanceRequest) SetImageURLBObject(v io.Reader) *DetectKitchenAnimalsAdvanceRequest {
	s.ImageURLBObject = v
	return s
}

type DetectKitchenAnimalsResponseBody struct {
	Data      *DetectKitchenAnimalsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectKitchenAnimalsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectKitchenAnimalsResponseBody) GoString() string {
	return s.String()
}

func (s *DetectKitchenAnimalsResponseBody) SetData(v *DetectKitchenAnimalsResponseBodyData) *DetectKitchenAnimalsResponseBody {
	s.Data = v
	return s
}

func (s *DetectKitchenAnimalsResponseBody) SetRequestId(v string) *DetectKitchenAnimalsResponseBody {
	s.RequestId = &v
	return s
}

type DetectKitchenAnimalsResponseBodyData struct {
	Elements []*DetectKitchenAnimalsResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s DetectKitchenAnimalsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectKitchenAnimalsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectKitchenAnimalsResponseBodyData) SetElements(v []*DetectKitchenAnimalsResponseBodyDataElements) *DetectKitchenAnimalsResponseBodyData {
	s.Elements = v
	return s
}

type DetectKitchenAnimalsResponseBodyDataElements struct {
	Rectangles *DetectKitchenAnimalsResponseBodyDataElementsRectangles `json:"Rectangles,omitempty" xml:"Rectangles,omitempty" type:"Struct"`
	Score      *float32                                                `json:"Score,omitempty" xml:"Score,omitempty"`
	Type       *string                                                 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DetectKitchenAnimalsResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s DetectKitchenAnimalsResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *DetectKitchenAnimalsResponseBodyDataElements) SetRectangles(v *DetectKitchenAnimalsResponseBodyDataElementsRectangles) *DetectKitchenAnimalsResponseBodyDataElements {
	s.Rectangles = v
	return s
}

func (s *DetectKitchenAnimalsResponseBodyDataElements) SetScore(v float32) *DetectKitchenAnimalsResponseBodyDataElements {
	s.Score = &v
	return s
}

func (s *DetectKitchenAnimalsResponseBodyDataElements) SetType(v string) *DetectKitchenAnimalsResponseBodyDataElements {
	s.Type = &v
	return s
}

type DetectKitchenAnimalsResponseBodyDataElementsRectangles struct {
	Height *int64 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int64 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int64 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int64 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s DetectKitchenAnimalsResponseBodyDataElementsRectangles) String() string {
	return tea.Prettify(s)
}

func (s DetectKitchenAnimalsResponseBodyDataElementsRectangles) GoString() string {
	return s.String()
}

func (s *DetectKitchenAnimalsResponseBodyDataElementsRectangles) SetHeight(v int64) *DetectKitchenAnimalsResponseBodyDataElementsRectangles {
	s.Height = &v
	return s
}

func (s *DetectKitchenAnimalsResponseBodyDataElementsRectangles) SetLeft(v int64) *DetectKitchenAnimalsResponseBodyDataElementsRectangles {
	s.Left = &v
	return s
}

func (s *DetectKitchenAnimalsResponseBodyDataElementsRectangles) SetTop(v int64) *DetectKitchenAnimalsResponseBodyDataElementsRectangles {
	s.Top = &v
	return s
}

func (s *DetectKitchenAnimalsResponseBodyDataElementsRectangles) SetWidth(v int64) *DetectKitchenAnimalsResponseBodyDataElementsRectangles {
	s.Width = &v
	return s
}

type DetectKitchenAnimalsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectKitchenAnimalsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectKitchenAnimalsResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectKitchenAnimalsResponse) GoString() string {
	return s.String()
}

func (s *DetectKitchenAnimalsResponse) SetHeaders(v map[string]*string) *DetectKitchenAnimalsResponse {
	s.Headers = v
	return s
}

func (s *DetectKitchenAnimalsResponse) SetStatusCode(v int32) *DetectKitchenAnimalsResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectKitchenAnimalsResponse) SetBody(v *DetectKitchenAnimalsResponseBody) *DetectKitchenAnimalsResponse {
	s.Body = v
	return s
}

type DetectMainBodyRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectMainBodyRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectMainBodyRequest) GoString() string {
	return s.String()
}

func (s *DetectMainBodyRequest) SetImageURL(v string) *DetectMainBodyRequest {
	s.ImageURL = &v
	return s
}

type DetectMainBodyAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectMainBodyAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectMainBodyAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectMainBodyAdvanceRequest) SetImageURLObject(v io.Reader) *DetectMainBodyAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type DetectMainBodyResponseBody struct {
	Data      *DetectMainBodyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectMainBodyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectMainBodyResponseBody) GoString() string {
	return s.String()
}

func (s *DetectMainBodyResponseBody) SetData(v *DetectMainBodyResponseBodyData) *DetectMainBodyResponseBody {
	s.Data = v
	return s
}

func (s *DetectMainBodyResponseBody) SetRequestId(v string) *DetectMainBodyResponseBody {
	s.RequestId = &v
	return s
}

type DetectMainBodyResponseBodyData struct {
	Location *DetectMainBodyResponseBodyDataLocation `json:"Location,omitempty" xml:"Location,omitempty" type:"Struct"`
}

func (s DetectMainBodyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectMainBodyResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectMainBodyResponseBodyData) SetLocation(v *DetectMainBodyResponseBodyDataLocation) *DetectMainBodyResponseBodyData {
	s.Location = v
	return s
}

type DetectMainBodyResponseBodyDataLocation struct {
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	X      *int32 `json:"X,omitempty" xml:"X,omitempty"`
	Y      *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DetectMainBodyResponseBodyDataLocation) String() string {
	return tea.Prettify(s)
}

func (s DetectMainBodyResponseBodyDataLocation) GoString() string {
	return s.String()
}

func (s *DetectMainBodyResponseBodyDataLocation) SetHeight(v int32) *DetectMainBodyResponseBodyDataLocation {
	s.Height = &v
	return s
}

func (s *DetectMainBodyResponseBodyDataLocation) SetWidth(v int32) *DetectMainBodyResponseBodyDataLocation {
	s.Width = &v
	return s
}

func (s *DetectMainBodyResponseBodyDataLocation) SetX(v int32) *DetectMainBodyResponseBodyDataLocation {
	s.X = &v
	return s
}

func (s *DetectMainBodyResponseBodyDataLocation) SetY(v int32) *DetectMainBodyResponseBodyDataLocation {
	s.Y = &v
	return s
}

type DetectMainBodyResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectMainBodyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectMainBodyResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectMainBodyResponse) GoString() string {
	return s.String()
}

func (s *DetectMainBodyResponse) SetHeaders(v map[string]*string) *DetectMainBodyResponse {
	s.Headers = v
	return s
}

func (s *DetectMainBodyResponse) SetStatusCode(v int32) *DetectMainBodyResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectMainBodyResponse) SetBody(v *DetectMainBodyResponseBody) *DetectMainBodyResponse {
	s.Body = v
	return s
}

type DetectObjectRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectObjectRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectObjectRequest) GoString() string {
	return s.String()
}

func (s *DetectObjectRequest) SetImageURL(v string) *DetectObjectRequest {
	s.ImageURL = &v
	return s
}

type DetectObjectAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectObjectAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectObjectAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectObjectAdvanceRequest) SetImageURLObject(v io.Reader) *DetectObjectAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type DetectObjectResponseBody struct {
	Data      *DetectObjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectObjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectObjectResponseBody) GoString() string {
	return s.String()
}

func (s *DetectObjectResponseBody) SetData(v *DetectObjectResponseBodyData) *DetectObjectResponseBody {
	s.Data = v
	return s
}

func (s *DetectObjectResponseBody) SetRequestId(v string) *DetectObjectResponseBody {
	s.RequestId = &v
	return s
}

type DetectObjectResponseBodyData struct {
	Elements []*DetectObjectResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	Height   *int32                                  `json:"Height,omitempty" xml:"Height,omitempty"`
	Width    *int32                                  `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s DetectObjectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectObjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectObjectResponseBodyData) SetElements(v []*DetectObjectResponseBodyDataElements) *DetectObjectResponseBodyData {
	s.Elements = v
	return s
}

func (s *DetectObjectResponseBodyData) SetHeight(v int32) *DetectObjectResponseBodyData {
	s.Height = &v
	return s
}

func (s *DetectObjectResponseBodyData) SetWidth(v int32) *DetectObjectResponseBodyData {
	s.Width = &v
	return s
}

type DetectObjectResponseBodyDataElements struct {
	Boxes []*int32 `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	Type  *string  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DetectObjectResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s DetectObjectResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *DetectObjectResponseBodyDataElements) SetBoxes(v []*int32) *DetectObjectResponseBodyDataElements {
	s.Boxes = v
	return s
}

func (s *DetectObjectResponseBodyDataElements) SetScore(v float32) *DetectObjectResponseBodyDataElements {
	s.Score = &v
	return s
}

func (s *DetectObjectResponseBodyDataElements) SetType(v string) *DetectObjectResponseBodyDataElements {
	s.Type = &v
	return s
}

type DetectObjectResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectObjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectObjectResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectObjectResponse) GoString() string {
	return s.String()
}

func (s *DetectObjectResponse) SetHeaders(v map[string]*string) *DetectObjectResponse {
	s.Headers = v
	return s
}

func (s *DetectObjectResponse) SetStatusCode(v int32) *DetectObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectObjectResponse) SetBody(v *DetectObjectResponseBody) *DetectObjectResponse {
	s.Body = v
	return s
}

type DetectTransparentImageRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectTransparentImageRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectTransparentImageRequest) GoString() string {
	return s.String()
}

func (s *DetectTransparentImageRequest) SetImageURL(v string) *DetectTransparentImageRequest {
	s.ImageURL = &v
	return s
}

type DetectTransparentImageAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectTransparentImageAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectTransparentImageAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectTransparentImageAdvanceRequest) SetImageURLObject(v io.Reader) *DetectTransparentImageAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type DetectTransparentImageResponseBody struct {
	Data      *DetectTransparentImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectTransparentImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectTransparentImageResponseBody) GoString() string {
	return s.String()
}

func (s *DetectTransparentImageResponseBody) SetData(v *DetectTransparentImageResponseBodyData) *DetectTransparentImageResponseBody {
	s.Data = v
	return s
}

func (s *DetectTransparentImageResponseBody) SetRequestId(v string) *DetectTransparentImageResponseBody {
	s.RequestId = &v
	return s
}

type DetectTransparentImageResponseBodyData struct {
	Elements []*DetectTransparentImageResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s DetectTransparentImageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectTransparentImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectTransparentImageResponseBodyData) SetElements(v []*DetectTransparentImageResponseBodyDataElements) *DetectTransparentImageResponseBodyData {
	s.Elements = v
	return s
}

type DetectTransparentImageResponseBodyDataElements struct {
	TransparentImage *int32 `json:"TransparentImage,omitempty" xml:"TransparentImage,omitempty"`
}

func (s DetectTransparentImageResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s DetectTransparentImageResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *DetectTransparentImageResponseBodyDataElements) SetTransparentImage(v int32) *DetectTransparentImageResponseBodyDataElements {
	s.TransparentImage = &v
	return s
}

type DetectTransparentImageResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectTransparentImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectTransparentImageResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectTransparentImageResponse) GoString() string {
	return s.String()
}

func (s *DetectTransparentImageResponse) SetHeaders(v map[string]*string) *DetectTransparentImageResponse {
	s.Headers = v
	return s
}

func (s *DetectTransparentImageResponse) SetStatusCode(v int32) *DetectTransparentImageResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectTransparentImageResponse) SetBody(v *DetectTransparentImageResponseBody) *DetectTransparentImageResponse {
	s.Body = v
	return s
}

type DetectVehicleRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectVehicleRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleRequest) GoString() string {
	return s.String()
}

func (s *DetectVehicleRequest) SetImageURL(v string) *DetectVehicleRequest {
	s.ImageURL = &v
	return s
}

type DetectVehicleAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectVehicleAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectVehicleAdvanceRequest) SetImageURLObject(v io.Reader) *DetectVehicleAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type DetectVehicleResponseBody struct {
	Data      *DetectVehicleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectVehicleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleResponseBody) GoString() string {
	return s.String()
}

func (s *DetectVehicleResponseBody) SetData(v *DetectVehicleResponseBodyData) *DetectVehicleResponseBody {
	s.Data = v
	return s
}

func (s *DetectVehicleResponseBody) SetRequestId(v string) *DetectVehicleResponseBody {
	s.RequestId = &v
	return s
}

type DetectVehicleResponseBodyData struct {
	DetectObjectInfoList []*DetectVehicleResponseBodyDataDetectObjectInfoList `json:"DetectObjectInfoList,omitempty" xml:"DetectObjectInfoList,omitempty" type:"Repeated"`
	Height               *int32                                               `json:"Height,omitempty" xml:"Height,omitempty"`
	Width                *int32                                               `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s DetectVehicleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectVehicleResponseBodyData) SetDetectObjectInfoList(v []*DetectVehicleResponseBodyDataDetectObjectInfoList) *DetectVehicleResponseBodyData {
	s.DetectObjectInfoList = v
	return s
}

func (s *DetectVehicleResponseBodyData) SetHeight(v int32) *DetectVehicleResponseBodyData {
	s.Height = &v
	return s
}

func (s *DetectVehicleResponseBodyData) SetWidth(v int32) *DetectVehicleResponseBodyData {
	s.Width = &v
	return s
}

type DetectVehicleResponseBodyDataDetectObjectInfoList struct {
	Boxes []*int32 `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
	Id    *int32   `json:"Id,omitempty" xml:"Id,omitempty"`
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	Type  *string  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DetectVehicleResponseBodyDataDetectObjectInfoList) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleResponseBodyDataDetectObjectInfoList) GoString() string {
	return s.String()
}

func (s *DetectVehicleResponseBodyDataDetectObjectInfoList) SetBoxes(v []*int32) *DetectVehicleResponseBodyDataDetectObjectInfoList {
	s.Boxes = v
	return s
}

func (s *DetectVehicleResponseBodyDataDetectObjectInfoList) SetId(v int32) *DetectVehicleResponseBodyDataDetectObjectInfoList {
	s.Id = &v
	return s
}

func (s *DetectVehicleResponseBodyDataDetectObjectInfoList) SetScore(v float32) *DetectVehicleResponseBodyDataDetectObjectInfoList {
	s.Score = &v
	return s
}

func (s *DetectVehicleResponseBodyDataDetectObjectInfoList) SetType(v string) *DetectVehicleResponseBodyDataDetectObjectInfoList {
	s.Type = &v
	return s
}

type DetectVehicleResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectVehicleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectVehicleResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleResponse) GoString() string {
	return s.String()
}

func (s *DetectVehicleResponse) SetHeaders(v map[string]*string) *DetectVehicleResponse {
	s.Headers = v
	return s
}

func (s *DetectVehicleResponse) SetStatusCode(v int32) *DetectVehicleResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectVehicleResponse) SetBody(v *DetectVehicleResponseBody) *DetectVehicleResponse {
	s.Body = v
	return s
}

type DetectVehicleICongestionRequest struct {
	ImageURL                   *string                                                      `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	PreRegionIntersectFeatures []*DetectVehicleICongestionRequestPreRegionIntersectFeatures `json:"PreRegionIntersectFeatures,omitempty" xml:"PreRegionIntersectFeatures,omitempty" type:"Repeated"`
	RoadRegions                []*DetectVehicleICongestionRequestRoadRegions                `json:"RoadRegions,omitempty" xml:"RoadRegions,omitempty" type:"Repeated"`
}

func (s DetectVehicleICongestionRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleICongestionRequest) GoString() string {
	return s.String()
}

func (s *DetectVehicleICongestionRequest) SetImageURL(v string) *DetectVehicleICongestionRequest {
	s.ImageURL = &v
	return s
}

func (s *DetectVehicleICongestionRequest) SetPreRegionIntersectFeatures(v []*DetectVehicleICongestionRequestPreRegionIntersectFeatures) *DetectVehicleICongestionRequest {
	s.PreRegionIntersectFeatures = v
	return s
}

func (s *DetectVehicleICongestionRequest) SetRoadRegions(v []*DetectVehicleICongestionRequestRoadRegions) *DetectVehicleICongestionRequest {
	s.RoadRegions = v
	return s
}

type DetectVehicleICongestionRequestPreRegionIntersectFeatures struct {
	Features []*string `json:"Features,omitempty" xml:"Features,omitempty" type:"Repeated"`
}

func (s DetectVehicleICongestionRequestPreRegionIntersectFeatures) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleICongestionRequestPreRegionIntersectFeatures) GoString() string {
	return s.String()
}

func (s *DetectVehicleICongestionRequestPreRegionIntersectFeatures) SetFeatures(v []*string) *DetectVehicleICongestionRequestPreRegionIntersectFeatures {
	s.Features = v
	return s
}

type DetectVehicleICongestionRequestRoadRegions struct {
	RoadRegion []*DetectVehicleICongestionRequestRoadRegionsRoadRegion `json:"RoadRegion,omitempty" xml:"RoadRegion,omitempty" type:"Repeated"`
}

func (s DetectVehicleICongestionRequestRoadRegions) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleICongestionRequestRoadRegions) GoString() string {
	return s.String()
}

func (s *DetectVehicleICongestionRequestRoadRegions) SetRoadRegion(v []*DetectVehicleICongestionRequestRoadRegionsRoadRegion) *DetectVehicleICongestionRequestRoadRegions {
	s.RoadRegion = v
	return s
}

type DetectVehicleICongestionRequestRoadRegionsRoadRegion struct {
	Point *DetectVehicleICongestionRequestRoadRegionsRoadRegionPoint `json:"Point,omitempty" xml:"Point,omitempty" type:"Struct"`
}

func (s DetectVehicleICongestionRequestRoadRegionsRoadRegion) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleICongestionRequestRoadRegionsRoadRegion) GoString() string {
	return s.String()
}

func (s *DetectVehicleICongestionRequestRoadRegionsRoadRegion) SetPoint(v *DetectVehicleICongestionRequestRoadRegionsRoadRegionPoint) *DetectVehicleICongestionRequestRoadRegionsRoadRegion {
	s.Point = v
	return s
}

type DetectVehicleICongestionRequestRoadRegionsRoadRegionPoint struct {
	X *int64 `json:"X,omitempty" xml:"X,omitempty"`
	Y *int64 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DetectVehicleICongestionRequestRoadRegionsRoadRegionPoint) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleICongestionRequestRoadRegionsRoadRegionPoint) GoString() string {
	return s.String()
}

func (s *DetectVehicleICongestionRequestRoadRegionsRoadRegionPoint) SetX(v int64) *DetectVehicleICongestionRequestRoadRegionsRoadRegionPoint {
	s.X = &v
	return s
}

func (s *DetectVehicleICongestionRequestRoadRegionsRoadRegionPoint) SetY(v int64) *DetectVehicleICongestionRequestRoadRegionsRoadRegionPoint {
	s.Y = &v
	return s
}

type DetectVehicleICongestionAdvanceRequest struct {
	ImageURLObject             io.Reader                                                           `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	PreRegionIntersectFeatures []*DetectVehicleICongestionAdvanceRequestPreRegionIntersectFeatures `json:"PreRegionIntersectFeatures,omitempty" xml:"PreRegionIntersectFeatures,omitempty" type:"Repeated"`
	RoadRegions                []*DetectVehicleICongestionAdvanceRequestRoadRegions                `json:"RoadRegions,omitempty" xml:"RoadRegions,omitempty" type:"Repeated"`
}

func (s DetectVehicleICongestionAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleICongestionAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectVehicleICongestionAdvanceRequest) SetImageURLObject(v io.Reader) *DetectVehicleICongestionAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *DetectVehicleICongestionAdvanceRequest) SetPreRegionIntersectFeatures(v []*DetectVehicleICongestionAdvanceRequestPreRegionIntersectFeatures) *DetectVehicleICongestionAdvanceRequest {
	s.PreRegionIntersectFeatures = v
	return s
}

func (s *DetectVehicleICongestionAdvanceRequest) SetRoadRegions(v []*DetectVehicleICongestionAdvanceRequestRoadRegions) *DetectVehicleICongestionAdvanceRequest {
	s.RoadRegions = v
	return s
}

type DetectVehicleICongestionAdvanceRequestPreRegionIntersectFeatures struct {
	Features []*string `json:"Features,omitempty" xml:"Features,omitempty" type:"Repeated"`
}

func (s DetectVehicleICongestionAdvanceRequestPreRegionIntersectFeatures) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleICongestionAdvanceRequestPreRegionIntersectFeatures) GoString() string {
	return s.String()
}

func (s *DetectVehicleICongestionAdvanceRequestPreRegionIntersectFeatures) SetFeatures(v []*string) *DetectVehicleICongestionAdvanceRequestPreRegionIntersectFeatures {
	s.Features = v
	return s
}

type DetectVehicleICongestionAdvanceRequestRoadRegions struct {
	RoadRegion []*DetectVehicleICongestionAdvanceRequestRoadRegionsRoadRegion `json:"RoadRegion,omitempty" xml:"RoadRegion,omitempty" type:"Repeated"`
}

func (s DetectVehicleICongestionAdvanceRequestRoadRegions) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleICongestionAdvanceRequestRoadRegions) GoString() string {
	return s.String()
}

func (s *DetectVehicleICongestionAdvanceRequestRoadRegions) SetRoadRegion(v []*DetectVehicleICongestionAdvanceRequestRoadRegionsRoadRegion) *DetectVehicleICongestionAdvanceRequestRoadRegions {
	s.RoadRegion = v
	return s
}

type DetectVehicleICongestionAdvanceRequestRoadRegionsRoadRegion struct {
	Point *DetectVehicleICongestionAdvanceRequestRoadRegionsRoadRegionPoint `json:"Point,omitempty" xml:"Point,omitempty" type:"Struct"`
}

func (s DetectVehicleICongestionAdvanceRequestRoadRegionsRoadRegion) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleICongestionAdvanceRequestRoadRegionsRoadRegion) GoString() string {
	return s.String()
}

func (s *DetectVehicleICongestionAdvanceRequestRoadRegionsRoadRegion) SetPoint(v *DetectVehicleICongestionAdvanceRequestRoadRegionsRoadRegionPoint) *DetectVehicleICongestionAdvanceRequestRoadRegionsRoadRegion {
	s.Point = v
	return s
}

type DetectVehicleICongestionAdvanceRequestRoadRegionsRoadRegionPoint struct {
	X *int64 `json:"X,omitempty" xml:"X,omitempty"`
	Y *int64 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DetectVehicleICongestionAdvanceRequestRoadRegionsRoadRegionPoint) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleICongestionAdvanceRequestRoadRegionsRoadRegionPoint) GoString() string {
	return s.String()
}

func (s *DetectVehicleICongestionAdvanceRequestRoadRegionsRoadRegionPoint) SetX(v int64) *DetectVehicleICongestionAdvanceRequestRoadRegionsRoadRegionPoint {
	s.X = &v
	return s
}

func (s *DetectVehicleICongestionAdvanceRequestRoadRegionsRoadRegionPoint) SetY(v int64) *DetectVehicleICongestionAdvanceRequestRoadRegionsRoadRegionPoint {
	s.Y = &v
	return s
}

type DetectVehicleICongestionShrinkRequest struct {
	ImageURL                         *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	PreRegionIntersectFeaturesShrink *string `json:"PreRegionIntersectFeatures,omitempty" xml:"PreRegionIntersectFeatures,omitempty"`
	RoadRegionsShrink                *string `json:"RoadRegions,omitempty" xml:"RoadRegions,omitempty"`
}

func (s DetectVehicleICongestionShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleICongestionShrinkRequest) GoString() string {
	return s.String()
}

func (s *DetectVehicleICongestionShrinkRequest) SetImageURL(v string) *DetectVehicleICongestionShrinkRequest {
	s.ImageURL = &v
	return s
}

func (s *DetectVehicleICongestionShrinkRequest) SetPreRegionIntersectFeaturesShrink(v string) *DetectVehicleICongestionShrinkRequest {
	s.PreRegionIntersectFeaturesShrink = &v
	return s
}

func (s *DetectVehicleICongestionShrinkRequest) SetRoadRegionsShrink(v string) *DetectVehicleICongestionShrinkRequest {
	s.RoadRegionsShrink = &v
	return s
}

type DetectVehicleICongestionResponseBody struct {
	Data      *DetectVehicleICongestionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectVehicleICongestionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleICongestionResponseBody) GoString() string {
	return s.String()
}

func (s *DetectVehicleICongestionResponseBody) SetData(v *DetectVehicleICongestionResponseBodyData) *DetectVehicleICongestionResponseBody {
	s.Data = v
	return s
}

func (s *DetectVehicleICongestionResponseBody) SetRequestId(v string) *DetectVehicleICongestionResponseBody {
	s.RequestId = &v
	return s
}

type DetectVehicleICongestionResponseBodyData struct {
	Elements                []*DetectVehicleICongestionResponseBodyDataElements                `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	RegionIntersectFeatures []*DetectVehicleICongestionResponseBodyDataRegionIntersectFeatures `json:"RegionIntersectFeatures,omitempty" xml:"RegionIntersectFeatures,omitempty" type:"Repeated"`
	RegionIntersectMatched  []*DetectVehicleICongestionResponseBodyDataRegionIntersectMatched  `json:"RegionIntersectMatched,omitempty" xml:"RegionIntersectMatched,omitempty" type:"Repeated"`
	RegionIntersects        []*DetectVehicleICongestionResponseBodyDataRegionIntersects        `json:"RegionIntersects,omitempty" xml:"RegionIntersects,omitempty" type:"Repeated"`
}

func (s DetectVehicleICongestionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleICongestionResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectVehicleICongestionResponseBodyData) SetElements(v []*DetectVehicleICongestionResponseBodyDataElements) *DetectVehicleICongestionResponseBodyData {
	s.Elements = v
	return s
}

func (s *DetectVehicleICongestionResponseBodyData) SetRegionIntersectFeatures(v []*DetectVehicleICongestionResponseBodyDataRegionIntersectFeatures) *DetectVehicleICongestionResponseBodyData {
	s.RegionIntersectFeatures = v
	return s
}

func (s *DetectVehicleICongestionResponseBodyData) SetRegionIntersectMatched(v []*DetectVehicleICongestionResponseBodyDataRegionIntersectMatched) *DetectVehicleICongestionResponseBodyData {
	s.RegionIntersectMatched = v
	return s
}

func (s *DetectVehicleICongestionResponseBodyData) SetRegionIntersects(v []*DetectVehicleICongestionResponseBodyDataRegionIntersects) *DetectVehicleICongestionResponseBodyData {
	s.RegionIntersects = v
	return s
}

type DetectVehicleICongestionResponseBodyDataElements struct {
	Boxes    []*DetectVehicleICongestionResponseBodyDataElementsBoxes `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
	Id       *int64                                                   `json:"Id,omitempty" xml:"Id,omitempty"`
	Score    *float32                                                 `json:"Score,omitempty" xml:"Score,omitempty"`
	TypeName *string                                                  `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
}

func (s DetectVehicleICongestionResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleICongestionResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *DetectVehicleICongestionResponseBodyDataElements) SetBoxes(v []*DetectVehicleICongestionResponseBodyDataElementsBoxes) *DetectVehicleICongestionResponseBodyDataElements {
	s.Boxes = v
	return s
}

func (s *DetectVehicleICongestionResponseBodyDataElements) SetId(v int64) *DetectVehicleICongestionResponseBodyDataElements {
	s.Id = &v
	return s
}

func (s *DetectVehicleICongestionResponseBodyDataElements) SetScore(v float32) *DetectVehicleICongestionResponseBodyDataElements {
	s.Score = &v
	return s
}

func (s *DetectVehicleICongestionResponseBodyDataElements) SetTypeName(v string) *DetectVehicleICongestionResponseBodyDataElements {
	s.TypeName = &v
	return s
}

type DetectVehicleICongestionResponseBodyDataElementsBoxes struct {
	Bottom *int64 `json:"Bottom,omitempty" xml:"Bottom,omitempty"`
	Left   *int64 `json:"Left,omitempty" xml:"Left,omitempty"`
	Right  *int64 `json:"Right,omitempty" xml:"Right,omitempty"`
	Top    *int64 `json:"Top,omitempty" xml:"Top,omitempty"`
}

func (s DetectVehicleICongestionResponseBodyDataElementsBoxes) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleICongestionResponseBodyDataElementsBoxes) GoString() string {
	return s.String()
}

func (s *DetectVehicleICongestionResponseBodyDataElementsBoxes) SetBottom(v int64) *DetectVehicleICongestionResponseBodyDataElementsBoxes {
	s.Bottom = &v
	return s
}

func (s *DetectVehicleICongestionResponseBodyDataElementsBoxes) SetLeft(v int64) *DetectVehicleICongestionResponseBodyDataElementsBoxes {
	s.Left = &v
	return s
}

func (s *DetectVehicleICongestionResponseBodyDataElementsBoxes) SetRight(v int64) *DetectVehicleICongestionResponseBodyDataElementsBoxes {
	s.Right = &v
	return s
}

func (s *DetectVehicleICongestionResponseBodyDataElementsBoxes) SetTop(v int64) *DetectVehicleICongestionResponseBodyDataElementsBoxes {
	s.Top = &v
	return s
}

type DetectVehicleICongestionResponseBodyDataRegionIntersectFeatures struct {
	Features []*string `json:"Features,omitempty" xml:"Features,omitempty" type:"Repeated"`
}

func (s DetectVehicleICongestionResponseBodyDataRegionIntersectFeatures) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleICongestionResponseBodyDataRegionIntersectFeatures) GoString() string {
	return s.String()
}

func (s *DetectVehicleICongestionResponseBodyDataRegionIntersectFeatures) SetFeatures(v []*string) *DetectVehicleICongestionResponseBodyDataRegionIntersectFeatures {
	s.Features = v
	return s
}

type DetectVehicleICongestionResponseBodyDataRegionIntersectMatched struct {
	Ids []*int64 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
}

func (s DetectVehicleICongestionResponseBodyDataRegionIntersectMatched) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleICongestionResponseBodyDataRegionIntersectMatched) GoString() string {
	return s.String()
}

func (s *DetectVehicleICongestionResponseBodyDataRegionIntersectMatched) SetIds(v []*int64) *DetectVehicleICongestionResponseBodyDataRegionIntersectMatched {
	s.Ids = v
	return s
}

type DetectVehicleICongestionResponseBodyDataRegionIntersects struct {
	Ids []*int64 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
}

func (s DetectVehicleICongestionResponseBodyDataRegionIntersects) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleICongestionResponseBodyDataRegionIntersects) GoString() string {
	return s.String()
}

func (s *DetectVehicleICongestionResponseBodyDataRegionIntersects) SetIds(v []*int64) *DetectVehicleICongestionResponseBodyDataRegionIntersects {
	s.Ids = v
	return s
}

type DetectVehicleICongestionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectVehicleICongestionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectVehicleICongestionResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleICongestionResponse) GoString() string {
	return s.String()
}

func (s *DetectVehicleICongestionResponse) SetHeaders(v map[string]*string) *DetectVehicleICongestionResponse {
	s.Headers = v
	return s
}

func (s *DetectVehicleICongestionResponse) SetStatusCode(v int32) *DetectVehicleICongestionResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectVehicleICongestionResponse) SetBody(v *DetectVehicleICongestionResponseBody) *DetectVehicleICongestionResponse {
	s.Body = v
	return s
}

type DetectVehicleIllegalParkingRequest struct {
	ImageURL    *string                                          `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	RoadRegions []*DetectVehicleIllegalParkingRequestRoadRegions `json:"RoadRegions,omitempty" xml:"RoadRegions,omitempty" type:"Repeated"`
}

func (s DetectVehicleIllegalParkingRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleIllegalParkingRequest) GoString() string {
	return s.String()
}

func (s *DetectVehicleIllegalParkingRequest) SetImageURL(v string) *DetectVehicleIllegalParkingRequest {
	s.ImageURL = &v
	return s
}

func (s *DetectVehicleIllegalParkingRequest) SetRoadRegions(v []*DetectVehicleIllegalParkingRequestRoadRegions) *DetectVehicleIllegalParkingRequest {
	s.RoadRegions = v
	return s
}

type DetectVehicleIllegalParkingRequestRoadRegions struct {
	RoadRegion []*DetectVehicleIllegalParkingRequestRoadRegionsRoadRegion `json:"RoadRegion,omitempty" xml:"RoadRegion,omitempty" type:"Repeated"`
}

func (s DetectVehicleIllegalParkingRequestRoadRegions) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleIllegalParkingRequestRoadRegions) GoString() string {
	return s.String()
}

func (s *DetectVehicleIllegalParkingRequestRoadRegions) SetRoadRegion(v []*DetectVehicleIllegalParkingRequestRoadRegionsRoadRegion) *DetectVehicleIllegalParkingRequestRoadRegions {
	s.RoadRegion = v
	return s
}

type DetectVehicleIllegalParkingRequestRoadRegionsRoadRegion struct {
	Point *DetectVehicleIllegalParkingRequestRoadRegionsRoadRegionPoint `json:"Point,omitempty" xml:"Point,omitempty" type:"Struct"`
}

func (s DetectVehicleIllegalParkingRequestRoadRegionsRoadRegion) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleIllegalParkingRequestRoadRegionsRoadRegion) GoString() string {
	return s.String()
}

func (s *DetectVehicleIllegalParkingRequestRoadRegionsRoadRegion) SetPoint(v *DetectVehicleIllegalParkingRequestRoadRegionsRoadRegionPoint) *DetectVehicleIllegalParkingRequestRoadRegionsRoadRegion {
	s.Point = v
	return s
}

type DetectVehicleIllegalParkingRequestRoadRegionsRoadRegionPoint struct {
	X *int64 `json:"X,omitempty" xml:"X,omitempty"`
	Y *int64 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DetectVehicleIllegalParkingRequestRoadRegionsRoadRegionPoint) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleIllegalParkingRequestRoadRegionsRoadRegionPoint) GoString() string {
	return s.String()
}

func (s *DetectVehicleIllegalParkingRequestRoadRegionsRoadRegionPoint) SetX(v int64) *DetectVehicleIllegalParkingRequestRoadRegionsRoadRegionPoint {
	s.X = &v
	return s
}

func (s *DetectVehicleIllegalParkingRequestRoadRegionsRoadRegionPoint) SetY(v int64) *DetectVehicleIllegalParkingRequestRoadRegionsRoadRegionPoint {
	s.Y = &v
	return s
}

type DetectVehicleIllegalParkingAdvanceRequest struct {
	ImageURLObject io.Reader                                               `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	RoadRegions    []*DetectVehicleIllegalParkingAdvanceRequestRoadRegions `json:"RoadRegions,omitempty" xml:"RoadRegions,omitempty" type:"Repeated"`
}

func (s DetectVehicleIllegalParkingAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleIllegalParkingAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectVehicleIllegalParkingAdvanceRequest) SetImageURLObject(v io.Reader) *DetectVehicleIllegalParkingAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *DetectVehicleIllegalParkingAdvanceRequest) SetRoadRegions(v []*DetectVehicleIllegalParkingAdvanceRequestRoadRegions) *DetectVehicleIllegalParkingAdvanceRequest {
	s.RoadRegions = v
	return s
}

type DetectVehicleIllegalParkingAdvanceRequestRoadRegions struct {
	RoadRegion []*DetectVehicleIllegalParkingAdvanceRequestRoadRegionsRoadRegion `json:"RoadRegion,omitempty" xml:"RoadRegion,omitempty" type:"Repeated"`
}

func (s DetectVehicleIllegalParkingAdvanceRequestRoadRegions) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleIllegalParkingAdvanceRequestRoadRegions) GoString() string {
	return s.String()
}

func (s *DetectVehicleIllegalParkingAdvanceRequestRoadRegions) SetRoadRegion(v []*DetectVehicleIllegalParkingAdvanceRequestRoadRegionsRoadRegion) *DetectVehicleIllegalParkingAdvanceRequestRoadRegions {
	s.RoadRegion = v
	return s
}

type DetectVehicleIllegalParkingAdvanceRequestRoadRegionsRoadRegion struct {
	Point *DetectVehicleIllegalParkingAdvanceRequestRoadRegionsRoadRegionPoint `json:"Point,omitempty" xml:"Point,omitempty" type:"Struct"`
}

func (s DetectVehicleIllegalParkingAdvanceRequestRoadRegionsRoadRegion) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleIllegalParkingAdvanceRequestRoadRegionsRoadRegion) GoString() string {
	return s.String()
}

func (s *DetectVehicleIllegalParkingAdvanceRequestRoadRegionsRoadRegion) SetPoint(v *DetectVehicleIllegalParkingAdvanceRequestRoadRegionsRoadRegionPoint) *DetectVehicleIllegalParkingAdvanceRequestRoadRegionsRoadRegion {
	s.Point = v
	return s
}

type DetectVehicleIllegalParkingAdvanceRequestRoadRegionsRoadRegionPoint struct {
	X *int64 `json:"X,omitempty" xml:"X,omitempty"`
	Y *int64 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DetectVehicleIllegalParkingAdvanceRequestRoadRegionsRoadRegionPoint) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleIllegalParkingAdvanceRequestRoadRegionsRoadRegionPoint) GoString() string {
	return s.String()
}

func (s *DetectVehicleIllegalParkingAdvanceRequestRoadRegionsRoadRegionPoint) SetX(v int64) *DetectVehicleIllegalParkingAdvanceRequestRoadRegionsRoadRegionPoint {
	s.X = &v
	return s
}

func (s *DetectVehicleIllegalParkingAdvanceRequestRoadRegionsRoadRegionPoint) SetY(v int64) *DetectVehicleIllegalParkingAdvanceRequestRoadRegionsRoadRegionPoint {
	s.Y = &v
	return s
}

type DetectVehicleIllegalParkingShrinkRequest struct {
	ImageURL          *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	RoadRegionsShrink *string `json:"RoadRegions,omitempty" xml:"RoadRegions,omitempty"`
}

func (s DetectVehicleIllegalParkingShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleIllegalParkingShrinkRequest) GoString() string {
	return s.String()
}

func (s *DetectVehicleIllegalParkingShrinkRequest) SetImageURL(v string) *DetectVehicleIllegalParkingShrinkRequest {
	s.ImageURL = &v
	return s
}

func (s *DetectVehicleIllegalParkingShrinkRequest) SetRoadRegionsShrink(v string) *DetectVehicleIllegalParkingShrinkRequest {
	s.RoadRegionsShrink = &v
	return s
}

type DetectVehicleIllegalParkingResponseBody struct {
	Data      *DetectVehicleIllegalParkingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectVehicleIllegalParkingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleIllegalParkingResponseBody) GoString() string {
	return s.String()
}

func (s *DetectVehicleIllegalParkingResponseBody) SetData(v *DetectVehicleIllegalParkingResponseBodyData) *DetectVehicleIllegalParkingResponseBody {
	s.Data = v
	return s
}

func (s *DetectVehicleIllegalParkingResponseBody) SetRequestId(v string) *DetectVehicleIllegalParkingResponseBody {
	s.RequestId = &v
	return s
}

type DetectVehicleIllegalParkingResponseBodyData struct {
	Elements         []*DetectVehicleIllegalParkingResponseBodyDataElements         `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	RegionIntersects []*DetectVehicleIllegalParkingResponseBodyDataRegionIntersects `json:"RegionIntersects,omitempty" xml:"RegionIntersects,omitempty" type:"Repeated"`
}

func (s DetectVehicleIllegalParkingResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleIllegalParkingResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectVehicleIllegalParkingResponseBodyData) SetElements(v []*DetectVehicleIllegalParkingResponseBodyDataElements) *DetectVehicleIllegalParkingResponseBodyData {
	s.Elements = v
	return s
}

func (s *DetectVehicleIllegalParkingResponseBodyData) SetRegionIntersects(v []*DetectVehicleIllegalParkingResponseBodyDataRegionIntersects) *DetectVehicleIllegalParkingResponseBodyData {
	s.RegionIntersects = v
	return s
}

type DetectVehicleIllegalParkingResponseBodyDataElements struct {
	Boxes    []*DetectVehicleIllegalParkingResponseBodyDataElementsBoxes `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
	Id       *int64                                                      `json:"Id,omitempty" xml:"Id,omitempty"`
	Score    *float32                                                    `json:"Score,omitempty" xml:"Score,omitempty"`
	TypeName *string                                                     `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
}

func (s DetectVehicleIllegalParkingResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleIllegalParkingResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *DetectVehicleIllegalParkingResponseBodyDataElements) SetBoxes(v []*DetectVehicleIllegalParkingResponseBodyDataElementsBoxes) *DetectVehicleIllegalParkingResponseBodyDataElements {
	s.Boxes = v
	return s
}

func (s *DetectVehicleIllegalParkingResponseBodyDataElements) SetId(v int64) *DetectVehicleIllegalParkingResponseBodyDataElements {
	s.Id = &v
	return s
}

func (s *DetectVehicleIllegalParkingResponseBodyDataElements) SetScore(v float32) *DetectVehicleIllegalParkingResponseBodyDataElements {
	s.Score = &v
	return s
}

func (s *DetectVehicleIllegalParkingResponseBodyDataElements) SetTypeName(v string) *DetectVehicleIllegalParkingResponseBodyDataElements {
	s.TypeName = &v
	return s
}

type DetectVehicleIllegalParkingResponseBodyDataElementsBoxes struct {
	Bottom *int64 `json:"Bottom,omitempty" xml:"Bottom,omitempty"`
	Left   *int64 `json:"Left,omitempty" xml:"Left,omitempty"`
	Right  *int64 `json:"Right,omitempty" xml:"Right,omitempty"`
	Top    *int64 `json:"Top,omitempty" xml:"Top,omitempty"`
}

func (s DetectVehicleIllegalParkingResponseBodyDataElementsBoxes) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleIllegalParkingResponseBodyDataElementsBoxes) GoString() string {
	return s.String()
}

func (s *DetectVehicleIllegalParkingResponseBodyDataElementsBoxes) SetBottom(v int64) *DetectVehicleIllegalParkingResponseBodyDataElementsBoxes {
	s.Bottom = &v
	return s
}

func (s *DetectVehicleIllegalParkingResponseBodyDataElementsBoxes) SetLeft(v int64) *DetectVehicleIllegalParkingResponseBodyDataElementsBoxes {
	s.Left = &v
	return s
}

func (s *DetectVehicleIllegalParkingResponseBodyDataElementsBoxes) SetRight(v int64) *DetectVehicleIllegalParkingResponseBodyDataElementsBoxes {
	s.Right = &v
	return s
}

func (s *DetectVehicleIllegalParkingResponseBodyDataElementsBoxes) SetTop(v int64) *DetectVehicleIllegalParkingResponseBodyDataElementsBoxes {
	s.Top = &v
	return s
}

type DetectVehicleIllegalParkingResponseBodyDataRegionIntersects struct {
	Ids []*int64 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
}

func (s DetectVehicleIllegalParkingResponseBodyDataRegionIntersects) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleIllegalParkingResponseBodyDataRegionIntersects) GoString() string {
	return s.String()
}

func (s *DetectVehicleIllegalParkingResponseBodyDataRegionIntersects) SetIds(v []*int64) *DetectVehicleIllegalParkingResponseBodyDataRegionIntersects {
	s.Ids = v
	return s
}

type DetectVehicleIllegalParkingResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectVehicleIllegalParkingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectVehicleIllegalParkingResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleIllegalParkingResponse) GoString() string {
	return s.String()
}

func (s *DetectVehicleIllegalParkingResponse) SetHeaders(v map[string]*string) *DetectVehicleIllegalParkingResponse {
	s.Headers = v
	return s
}

func (s *DetectVehicleIllegalParkingResponse) SetStatusCode(v int32) *DetectVehicleIllegalParkingResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectVehicleIllegalParkingResponse) SetBody(v *DetectVehicleIllegalParkingResponseBody) *DetectVehicleIllegalParkingResponse {
	s.Body = v
	return s
}

type DetectVideoIPCObjectRequest struct {
	CallbackOnlyHasObject *bool   `json:"CallbackOnlyHasObject,omitempty" xml:"CallbackOnlyHasObject,omitempty"`
	StartTimestamp        *int64  `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	VideoURL              *string `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
}

func (s DetectVideoIPCObjectRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectVideoIPCObjectRequest) GoString() string {
	return s.String()
}

func (s *DetectVideoIPCObjectRequest) SetCallbackOnlyHasObject(v bool) *DetectVideoIPCObjectRequest {
	s.CallbackOnlyHasObject = &v
	return s
}

func (s *DetectVideoIPCObjectRequest) SetStartTimestamp(v int64) *DetectVideoIPCObjectRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DetectVideoIPCObjectRequest) SetVideoURL(v string) *DetectVideoIPCObjectRequest {
	s.VideoURL = &v
	return s
}

type DetectVideoIPCObjectAdvanceRequest struct {
	CallbackOnlyHasObject *bool     `json:"CallbackOnlyHasObject,omitempty" xml:"CallbackOnlyHasObject,omitempty"`
	StartTimestamp        *int64    `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	VideoURLObject        io.Reader `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
}

func (s DetectVideoIPCObjectAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectVideoIPCObjectAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectVideoIPCObjectAdvanceRequest) SetCallbackOnlyHasObject(v bool) *DetectVideoIPCObjectAdvanceRequest {
	s.CallbackOnlyHasObject = &v
	return s
}

func (s *DetectVideoIPCObjectAdvanceRequest) SetStartTimestamp(v int64) *DetectVideoIPCObjectAdvanceRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DetectVideoIPCObjectAdvanceRequest) SetVideoURLObject(v io.Reader) *DetectVideoIPCObjectAdvanceRequest {
	s.VideoURLObject = v
	return s
}

type DetectVideoIPCObjectResponseBody struct {
	Data      *DetectVideoIPCObjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectVideoIPCObjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectVideoIPCObjectResponseBody) GoString() string {
	return s.String()
}

func (s *DetectVideoIPCObjectResponseBody) SetData(v *DetectVideoIPCObjectResponseBodyData) *DetectVideoIPCObjectResponseBody {
	s.Data = v
	return s
}

func (s *DetectVideoIPCObjectResponseBody) SetMessage(v string) *DetectVideoIPCObjectResponseBody {
	s.Message = &v
	return s
}

func (s *DetectVideoIPCObjectResponseBody) SetRequestId(v string) *DetectVideoIPCObjectResponseBody {
	s.RequestId = &v
	return s
}

type DetectVideoIPCObjectResponseBodyData struct {
	Frames    []*DetectVideoIPCObjectResponseBodyDataFrames `json:"Frames,omitempty" xml:"Frames,omitempty" type:"Repeated"`
	Height    *int64                                        `json:"Height,omitempty" xml:"Height,omitempty"`
	InputFile *string                                       `json:"InputFile,omitempty" xml:"InputFile,omitempty"`
	Width     *int64                                        `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s DetectVideoIPCObjectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectVideoIPCObjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectVideoIPCObjectResponseBodyData) SetFrames(v []*DetectVideoIPCObjectResponseBodyDataFrames) *DetectVideoIPCObjectResponseBodyData {
	s.Frames = v
	return s
}

func (s *DetectVideoIPCObjectResponseBodyData) SetHeight(v int64) *DetectVideoIPCObjectResponseBodyData {
	s.Height = &v
	return s
}

func (s *DetectVideoIPCObjectResponseBodyData) SetInputFile(v string) *DetectVideoIPCObjectResponseBodyData {
	s.InputFile = &v
	return s
}

func (s *DetectVideoIPCObjectResponseBodyData) SetWidth(v int64) *DetectVideoIPCObjectResponseBodyData {
	s.Width = &v
	return s
}

type DetectVideoIPCObjectResponseBodyDataFrames struct {
	Elements []*DetectVideoIPCObjectResponseBodyDataFramesElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	Time     *int64                                                `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DetectVideoIPCObjectResponseBodyDataFrames) String() string {
	return tea.Prettify(s)
}

func (s DetectVideoIPCObjectResponseBodyDataFrames) GoString() string {
	return s.String()
}

func (s *DetectVideoIPCObjectResponseBodyDataFrames) SetElements(v []*DetectVideoIPCObjectResponseBodyDataFramesElements) *DetectVideoIPCObjectResponseBodyDataFrames {
	s.Elements = v
	return s
}

func (s *DetectVideoIPCObjectResponseBodyDataFrames) SetTime(v int64) *DetectVideoIPCObjectResponseBodyDataFrames {
	s.Time = &v
	return s
}

type DetectVideoIPCObjectResponseBodyDataFramesElements struct {
	Height *int64   `json:"Height,omitempty" xml:"Height,omitempty"`
	Score  *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	Type   *string  `json:"Type,omitempty" xml:"Type,omitempty"`
	Width  *int64   `json:"Width,omitempty" xml:"Width,omitempty"`
	X      *int64   `json:"X,omitempty" xml:"X,omitempty"`
	Y      *int64   `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DetectVideoIPCObjectResponseBodyDataFramesElements) String() string {
	return tea.Prettify(s)
}

func (s DetectVideoIPCObjectResponseBodyDataFramesElements) GoString() string {
	return s.String()
}

func (s *DetectVideoIPCObjectResponseBodyDataFramesElements) SetHeight(v int64) *DetectVideoIPCObjectResponseBodyDataFramesElements {
	s.Height = &v
	return s
}

func (s *DetectVideoIPCObjectResponseBodyDataFramesElements) SetScore(v float32) *DetectVideoIPCObjectResponseBodyDataFramesElements {
	s.Score = &v
	return s
}

func (s *DetectVideoIPCObjectResponseBodyDataFramesElements) SetType(v string) *DetectVideoIPCObjectResponseBodyDataFramesElements {
	s.Type = &v
	return s
}

func (s *DetectVideoIPCObjectResponseBodyDataFramesElements) SetWidth(v int64) *DetectVideoIPCObjectResponseBodyDataFramesElements {
	s.Width = &v
	return s
}

func (s *DetectVideoIPCObjectResponseBodyDataFramesElements) SetX(v int64) *DetectVideoIPCObjectResponseBodyDataFramesElements {
	s.X = &v
	return s
}

func (s *DetectVideoIPCObjectResponseBodyDataFramesElements) SetY(v int64) *DetectVideoIPCObjectResponseBodyDataFramesElements {
	s.Y = &v
	return s
}

type DetectVideoIPCObjectResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectVideoIPCObjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectVideoIPCObjectResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectVideoIPCObjectResponse) GoString() string {
	return s.String()
}

func (s *DetectVideoIPCObjectResponse) SetHeaders(v map[string]*string) *DetectVideoIPCObjectResponse {
	s.Headers = v
	return s
}

func (s *DetectVideoIPCObjectResponse) SetStatusCode(v int32) *DetectVideoIPCObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectVideoIPCObjectResponse) SetBody(v *DetectVideoIPCObjectResponseBody) *DetectVideoIPCObjectResponse {
	s.Body = v
	return s
}

type DetectWhiteBaseImageRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectWhiteBaseImageRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectWhiteBaseImageRequest) GoString() string {
	return s.String()
}

func (s *DetectWhiteBaseImageRequest) SetImageURL(v string) *DetectWhiteBaseImageRequest {
	s.ImageURL = &v
	return s
}

type DetectWhiteBaseImageAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectWhiteBaseImageAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectWhiteBaseImageAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectWhiteBaseImageAdvanceRequest) SetImageURLObject(v io.Reader) *DetectWhiteBaseImageAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type DetectWhiteBaseImageResponseBody struct {
	Data      *DetectWhiteBaseImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectWhiteBaseImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectWhiteBaseImageResponseBody) GoString() string {
	return s.String()
}

func (s *DetectWhiteBaseImageResponseBody) SetData(v *DetectWhiteBaseImageResponseBodyData) *DetectWhiteBaseImageResponseBody {
	s.Data = v
	return s
}

func (s *DetectWhiteBaseImageResponseBody) SetRequestId(v string) *DetectWhiteBaseImageResponseBody {
	s.RequestId = &v
	return s
}

type DetectWhiteBaseImageResponseBodyData struct {
	Elements []*DetectWhiteBaseImageResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s DetectWhiteBaseImageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectWhiteBaseImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectWhiteBaseImageResponseBodyData) SetElements(v []*DetectWhiteBaseImageResponseBodyDataElements) *DetectWhiteBaseImageResponseBodyData {
	s.Elements = v
	return s
}

type DetectWhiteBaseImageResponseBodyDataElements struct {
	WhiteBase *int32 `json:"WhiteBase,omitempty" xml:"WhiteBase,omitempty"`
}

func (s DetectWhiteBaseImageResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s DetectWhiteBaseImageResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *DetectWhiteBaseImageResponseBodyDataElements) SetWhiteBase(v int32) *DetectWhiteBaseImageResponseBodyDataElements {
	s.WhiteBase = &v
	return s
}

type DetectWhiteBaseImageResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectWhiteBaseImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectWhiteBaseImageResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectWhiteBaseImageResponse) GoString() string {
	return s.String()
}

func (s *DetectWhiteBaseImageResponse) SetHeaders(v map[string]*string) *DetectWhiteBaseImageResponse {
	s.Headers = v
	return s
}

func (s *DetectWhiteBaseImageResponse) SetStatusCode(v int32) *DetectWhiteBaseImageResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectWhiteBaseImageResponse) SetBody(v *DetectWhiteBaseImageResponseBody) *DetectWhiteBaseImageResponse {
	s.Body = v
	return s
}

type DetectWorkwearRequest struct {
	Clothes  *DetectWorkwearRequestClothes `json:"Clothes,omitempty" xml:"Clothes,omitempty" type:"Struct"`
	ImageUrl *string                       `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// 1
	Labels []*string `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
}

func (s DetectWorkwearRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectWorkwearRequest) GoString() string {
	return s.String()
}

func (s *DetectWorkwearRequest) SetClothes(v *DetectWorkwearRequestClothes) *DetectWorkwearRequest {
	s.Clothes = v
	return s
}

func (s *DetectWorkwearRequest) SetImageUrl(v string) *DetectWorkwearRequest {
	s.ImageUrl = &v
	return s
}

func (s *DetectWorkwearRequest) SetLabels(v []*string) *DetectWorkwearRequest {
	s.Labels = v
	return s
}

type DetectWorkwearRequestClothes struct {
	MaxNum    *int64   `json:"MaxNum,omitempty" xml:"MaxNum,omitempty"`
	Threshold *float64 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s DetectWorkwearRequestClothes) String() string {
	return tea.Prettify(s)
}

func (s DetectWorkwearRequestClothes) GoString() string {
	return s.String()
}

func (s *DetectWorkwearRequestClothes) SetMaxNum(v int64) *DetectWorkwearRequestClothes {
	s.MaxNum = &v
	return s
}

func (s *DetectWorkwearRequestClothes) SetThreshold(v float64) *DetectWorkwearRequestClothes {
	s.Threshold = &v
	return s
}

type DetectWorkwearAdvanceRequest struct {
	Clothes        *DetectWorkwearAdvanceRequestClothes `json:"Clothes,omitempty" xml:"Clothes,omitempty" type:"Struct"`
	ImageUrlObject io.Reader                            `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// 1
	Labels []*string `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
}

func (s DetectWorkwearAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectWorkwearAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectWorkwearAdvanceRequest) SetClothes(v *DetectWorkwearAdvanceRequestClothes) *DetectWorkwearAdvanceRequest {
	s.Clothes = v
	return s
}

func (s *DetectWorkwearAdvanceRequest) SetImageUrlObject(v io.Reader) *DetectWorkwearAdvanceRequest {
	s.ImageUrlObject = v
	return s
}

func (s *DetectWorkwearAdvanceRequest) SetLabels(v []*string) *DetectWorkwearAdvanceRequest {
	s.Labels = v
	return s
}

type DetectWorkwearAdvanceRequestClothes struct {
	MaxNum    *int64   `json:"MaxNum,omitempty" xml:"MaxNum,omitempty"`
	Threshold *float64 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s DetectWorkwearAdvanceRequestClothes) String() string {
	return tea.Prettify(s)
}

func (s DetectWorkwearAdvanceRequestClothes) GoString() string {
	return s.String()
}

func (s *DetectWorkwearAdvanceRequestClothes) SetMaxNum(v int64) *DetectWorkwearAdvanceRequestClothes {
	s.MaxNum = &v
	return s
}

func (s *DetectWorkwearAdvanceRequestClothes) SetThreshold(v float64) *DetectWorkwearAdvanceRequestClothes {
	s.Threshold = &v
	return s
}

type DetectWorkwearShrinkRequest struct {
	ClothesShrink *string `json:"Clothes,omitempty" xml:"Clothes,omitempty"`
	ImageUrl      *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// 1
	Labels []*string `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
}

func (s DetectWorkwearShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectWorkwearShrinkRequest) GoString() string {
	return s.String()
}

func (s *DetectWorkwearShrinkRequest) SetClothesShrink(v string) *DetectWorkwearShrinkRequest {
	s.ClothesShrink = &v
	return s
}

func (s *DetectWorkwearShrinkRequest) SetImageUrl(v string) *DetectWorkwearShrinkRequest {
	s.ImageUrl = &v
	return s
}

func (s *DetectWorkwearShrinkRequest) SetLabels(v []*string) *DetectWorkwearShrinkRequest {
	s.Labels = v
	return s
}

type DetectWorkwearResponseBody struct {
	Data      *DetectWorkwearResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectWorkwearResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectWorkwearResponseBody) GoString() string {
	return s.String()
}

func (s *DetectWorkwearResponseBody) SetData(v *DetectWorkwearResponseBodyData) *DetectWorkwearResponseBody {
	s.Data = v
	return s
}

func (s *DetectWorkwearResponseBody) SetRequestId(v string) *DetectWorkwearResponseBody {
	s.RequestId = &v
	return s
}

type DetectWorkwearResponseBodyData struct {
	Elements []*DetectWorkwearResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s DetectWorkwearResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectWorkwearResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectWorkwearResponseBodyData) SetElements(v []*DetectWorkwearResponseBodyDataElements) *DetectWorkwearResponseBodyData {
	s.Elements = v
	return s
}

type DetectWorkwearResponseBodyDataElements struct {
	Property   []*DetectWorkwearResponseBodyDataElementsProperty `json:"Property,omitempty" xml:"Property,omitempty" type:"Repeated"`
	Rectangles *DetectWorkwearResponseBodyDataElementsRectangles `json:"Rectangles,omitempty" xml:"Rectangles,omitempty" type:"Struct"`
	Score      *float64                                          `json:"Score,omitempty" xml:"Score,omitempty"`
	Type       *string                                           `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DetectWorkwearResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s DetectWorkwearResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *DetectWorkwearResponseBodyDataElements) SetProperty(v []*DetectWorkwearResponseBodyDataElementsProperty) *DetectWorkwearResponseBodyDataElements {
	s.Property = v
	return s
}

func (s *DetectWorkwearResponseBodyDataElements) SetRectangles(v *DetectWorkwearResponseBodyDataElementsRectangles) *DetectWorkwearResponseBodyDataElements {
	s.Rectangles = v
	return s
}

func (s *DetectWorkwearResponseBodyDataElements) SetScore(v float64) *DetectWorkwearResponseBodyDataElements {
	s.Score = &v
	return s
}

func (s *DetectWorkwearResponseBodyDataElements) SetType(v string) *DetectWorkwearResponseBodyDataElements {
	s.Type = &v
	return s
}

type DetectWorkwearResponseBodyDataElementsProperty struct {
	Label       *string                                                    `json:"Label,omitempty" xml:"Label,omitempty"`
	Probability *DetectWorkwearResponseBodyDataElementsPropertyProbability `json:"Probability,omitempty" xml:"Probability,omitempty" type:"Struct"`
}

func (s DetectWorkwearResponseBodyDataElementsProperty) String() string {
	return tea.Prettify(s)
}

func (s DetectWorkwearResponseBodyDataElementsProperty) GoString() string {
	return s.String()
}

func (s *DetectWorkwearResponseBodyDataElementsProperty) SetLabel(v string) *DetectWorkwearResponseBodyDataElementsProperty {
	s.Label = &v
	return s
}

func (s *DetectWorkwearResponseBodyDataElementsProperty) SetProbability(v *DetectWorkwearResponseBodyDataElementsPropertyProbability) *DetectWorkwearResponseBodyDataElementsProperty {
	s.Probability = v
	return s
}

type DetectWorkwearResponseBodyDataElementsPropertyProbability struct {
	No        *float64 `json:"No,omitempty" xml:"No,omitempty"`
	Threshold *int64   `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Unknown   *float64 `json:"Unknown,omitempty" xml:"Unknown,omitempty"`
	Yes       *float64 `json:"Yes,omitempty" xml:"Yes,omitempty"`
}

func (s DetectWorkwearResponseBodyDataElementsPropertyProbability) String() string {
	return tea.Prettify(s)
}

func (s DetectWorkwearResponseBodyDataElementsPropertyProbability) GoString() string {
	return s.String()
}

func (s *DetectWorkwearResponseBodyDataElementsPropertyProbability) SetNo(v float64) *DetectWorkwearResponseBodyDataElementsPropertyProbability {
	s.No = &v
	return s
}

func (s *DetectWorkwearResponseBodyDataElementsPropertyProbability) SetThreshold(v int64) *DetectWorkwearResponseBodyDataElementsPropertyProbability {
	s.Threshold = &v
	return s
}

func (s *DetectWorkwearResponseBodyDataElementsPropertyProbability) SetUnknown(v float64) *DetectWorkwearResponseBodyDataElementsPropertyProbability {
	s.Unknown = &v
	return s
}

func (s *DetectWorkwearResponseBodyDataElementsPropertyProbability) SetYes(v float64) *DetectWorkwearResponseBodyDataElementsPropertyProbability {
	s.Yes = &v
	return s
}

type DetectWorkwearResponseBodyDataElementsRectangles struct {
	Height *int64 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int64 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int64 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int64 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s DetectWorkwearResponseBodyDataElementsRectangles) String() string {
	return tea.Prettify(s)
}

func (s DetectWorkwearResponseBodyDataElementsRectangles) GoString() string {
	return s.String()
}

func (s *DetectWorkwearResponseBodyDataElementsRectangles) SetHeight(v int64) *DetectWorkwearResponseBodyDataElementsRectangles {
	s.Height = &v
	return s
}

func (s *DetectWorkwearResponseBodyDataElementsRectangles) SetLeft(v int64) *DetectWorkwearResponseBodyDataElementsRectangles {
	s.Left = &v
	return s
}

func (s *DetectWorkwearResponseBodyDataElementsRectangles) SetTop(v int64) *DetectWorkwearResponseBodyDataElementsRectangles {
	s.Top = &v
	return s
}

func (s *DetectWorkwearResponseBodyDataElementsRectangles) SetWidth(v int64) *DetectWorkwearResponseBodyDataElementsRectangles {
	s.Width = &v
	return s
}

type DetectWorkwearResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectWorkwearResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectWorkwearResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectWorkwearResponse) GoString() string {
	return s.String()
}

func (s *DetectWorkwearResponse) SetHeaders(v map[string]*string) *DetectWorkwearResponse {
	s.Headers = v
	return s
}

func (s *DetectWorkwearResponse) SetStatusCode(v int32) *DetectWorkwearResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectWorkwearResponse) SetBody(v *DetectWorkwearResponseBody) *DetectWorkwearResponse {
	s.Body = v
	return s
}

type GetAsyncJobResultRequest struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetAsyncJobResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultRequest) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultRequest) SetJobId(v string) *GetAsyncJobResultRequest {
	s.JobId = &v
	return s
}

type GetAsyncJobResultResponseBody struct {
	Data      *GetAsyncJobResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAsyncJobResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultResponseBody) SetData(v *GetAsyncJobResultResponseBodyData) *GetAsyncJobResultResponseBody {
	s.Data = v
	return s
}

func (s *GetAsyncJobResultResponseBody) SetRequestId(v string) *GetAsyncJobResultResponseBody {
	s.RequestId = &v
	return s
}

type GetAsyncJobResultResponseBodyData struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Result       *string `json:"Result,omitempty" xml:"Result,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAsyncJobResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultResponseBodyData) SetErrorCode(v string) *GetAsyncJobResultResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetErrorMessage(v string) *GetAsyncJobResultResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetJobId(v string) *GetAsyncJobResultResponseBodyData {
	s.JobId = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetResult(v string) *GetAsyncJobResultResponseBodyData {
	s.Result = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetStatus(v string) *GetAsyncJobResultResponseBodyData {
	s.Status = &v
	return s
}

type GetAsyncJobResultResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAsyncJobResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAsyncJobResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultResponse) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultResponse) SetHeaders(v map[string]*string) *GetAsyncJobResultResponse {
	s.Headers = v
	return s
}

func (s *GetAsyncJobResultResponse) SetStatusCode(v int32) *GetAsyncJobResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAsyncJobResultResponse) SetBody(v *GetAsyncJobResultResponseBody) *GetAsyncJobResultResponse {
	s.Body = v
	return s
}

type RecognizeVehicleDamageRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeVehicleDamageRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehicleDamageRequest) GoString() string {
	return s.String()
}

func (s *RecognizeVehicleDamageRequest) SetImageURL(v string) *RecognizeVehicleDamageRequest {
	s.ImageURL = &v
	return s
}

type RecognizeVehicleDamageAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeVehicleDamageAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehicleDamageAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeVehicleDamageAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeVehicleDamageAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type RecognizeVehicleDamageResponseBody struct {
	Data      *RecognizeVehicleDamageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeVehicleDamageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehicleDamageResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeVehicleDamageResponseBody) SetData(v *RecognizeVehicleDamageResponseBodyData) *RecognizeVehicleDamageResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeVehicleDamageResponseBody) SetRequestId(v string) *RecognizeVehicleDamageResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeVehicleDamageResponseBodyData struct {
	Elements []*RecognizeVehicleDamageResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s RecognizeVehicleDamageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehicleDamageResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeVehicleDamageResponseBodyData) SetElements(v []*RecognizeVehicleDamageResponseBodyDataElements) *RecognizeVehicleDamageResponseBodyData {
	s.Elements = v
	return s
}

type RecognizeVehicleDamageResponseBodyDataElements struct {
	// 1
	Boxes []*int32 `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// 1
	Scores []*float32 `json:"Scores,omitempty" xml:"Scores,omitempty" type:"Repeated"`
	Type   *string    `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s RecognizeVehicleDamageResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehicleDamageResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *RecognizeVehicleDamageResponseBodyDataElements) SetBoxes(v []*int32) *RecognizeVehicleDamageResponseBodyDataElements {
	s.Boxes = v
	return s
}

func (s *RecognizeVehicleDamageResponseBodyDataElements) SetScore(v float32) *RecognizeVehicleDamageResponseBodyDataElements {
	s.Score = &v
	return s
}

func (s *RecognizeVehicleDamageResponseBodyDataElements) SetScores(v []*float32) *RecognizeVehicleDamageResponseBodyDataElements {
	s.Scores = v
	return s
}

func (s *RecognizeVehicleDamageResponseBodyDataElements) SetType(v string) *RecognizeVehicleDamageResponseBodyDataElements {
	s.Type = &v
	return s
}

type RecognizeVehicleDamageResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeVehicleDamageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeVehicleDamageResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehicleDamageResponse) GoString() string {
	return s.String()
}

func (s *RecognizeVehicleDamageResponse) SetHeaders(v map[string]*string) *RecognizeVehicleDamageResponse {
	s.Headers = v
	return s
}

func (s *RecognizeVehicleDamageResponse) SetStatusCode(v int32) *RecognizeVehicleDamageResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeVehicleDamageResponse) SetBody(v *RecognizeVehicleDamageResponseBody) *RecognizeVehicleDamageResponse {
	s.Body = v
	return s
}

type RecognizeVehicleDashboardRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeVehicleDashboardRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehicleDashboardRequest) GoString() string {
	return s.String()
}

func (s *RecognizeVehicleDashboardRequest) SetImageURL(v string) *RecognizeVehicleDashboardRequest {
	s.ImageURL = &v
	return s
}

type RecognizeVehicleDashboardAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeVehicleDashboardAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehicleDashboardAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeVehicleDashboardAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeVehicleDashboardAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type RecognizeVehicleDashboardResponseBody struct {
	Data      *RecognizeVehicleDashboardResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeVehicleDashboardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehicleDashboardResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeVehicleDashboardResponseBody) SetData(v *RecognizeVehicleDashboardResponseBodyData) *RecognizeVehicleDashboardResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeVehicleDashboardResponseBody) SetRequestId(v string) *RecognizeVehicleDashboardResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeVehicleDashboardResponseBodyData struct {
	Elements []*RecognizeVehicleDashboardResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s RecognizeVehicleDashboardResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehicleDashboardResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeVehicleDashboardResponseBodyData) SetElements(v []*RecognizeVehicleDashboardResponseBodyDataElements) *RecognizeVehicleDashboardResponseBodyData {
	s.Elements = v
	return s
}

type RecognizeVehicleDashboardResponseBodyDataElements struct {
	Boxes     []*float32 `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
	ClassName *string    `json:"ClassName,omitempty" xml:"ClassName,omitempty"`
	Label     *string    `json:"Label,omitempty" xml:"Label,omitempty"`
	Score     *float32   `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s RecognizeVehicleDashboardResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehicleDashboardResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *RecognizeVehicleDashboardResponseBodyDataElements) SetBoxes(v []*float32) *RecognizeVehicleDashboardResponseBodyDataElements {
	s.Boxes = v
	return s
}

func (s *RecognizeVehicleDashboardResponseBodyDataElements) SetClassName(v string) *RecognizeVehicleDashboardResponseBodyDataElements {
	s.ClassName = &v
	return s
}

func (s *RecognizeVehicleDashboardResponseBodyDataElements) SetLabel(v string) *RecognizeVehicleDashboardResponseBodyDataElements {
	s.Label = &v
	return s
}

func (s *RecognizeVehicleDashboardResponseBodyDataElements) SetScore(v float32) *RecognizeVehicleDashboardResponseBodyDataElements {
	s.Score = &v
	return s
}

type RecognizeVehicleDashboardResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeVehicleDashboardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeVehicleDashboardResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehicleDashboardResponse) GoString() string {
	return s.String()
}

func (s *RecognizeVehicleDashboardResponse) SetHeaders(v map[string]*string) *RecognizeVehicleDashboardResponse {
	s.Headers = v
	return s
}

func (s *RecognizeVehicleDashboardResponse) SetStatusCode(v int32) *RecognizeVehicleDashboardResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeVehicleDashboardResponse) SetBody(v *RecognizeVehicleDashboardResponseBody) *RecognizeVehicleDashboardResponse {
	s.Body = v
	return s
}

type RecognizeVehiclePartsRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeVehiclePartsRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehiclePartsRequest) GoString() string {
	return s.String()
}

func (s *RecognizeVehiclePartsRequest) SetImageURL(v string) *RecognizeVehiclePartsRequest {
	s.ImageURL = &v
	return s
}

type RecognizeVehiclePartsAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeVehiclePartsAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehiclePartsAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeVehiclePartsAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeVehiclePartsAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type RecognizeVehiclePartsResponseBody struct {
	Data      *RecognizeVehiclePartsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeVehiclePartsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehiclePartsResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeVehiclePartsResponseBody) SetData(v *RecognizeVehiclePartsResponseBodyData) *RecognizeVehiclePartsResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeVehiclePartsResponseBody) SetRequestId(v string) *RecognizeVehiclePartsResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeVehiclePartsResponseBodyData struct {
	Elements []*RecognizeVehiclePartsResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	// 1
	OriginShapes []*int32 `json:"OriginShapes,omitempty" xml:"OriginShapes,omitempty" type:"Repeated"`
}

func (s RecognizeVehiclePartsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehiclePartsResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeVehiclePartsResponseBodyData) SetElements(v []*RecognizeVehiclePartsResponseBodyDataElements) *RecognizeVehiclePartsResponseBodyData {
	s.Elements = v
	return s
}

func (s *RecognizeVehiclePartsResponseBodyData) SetOriginShapes(v []*int32) *RecognizeVehiclePartsResponseBodyData {
	s.OriginShapes = v
	return s
}

type RecognizeVehiclePartsResponseBodyDataElements struct {
	// 1
	Boxes []*int32 `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	Type  *string  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s RecognizeVehiclePartsResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehiclePartsResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *RecognizeVehiclePartsResponseBodyDataElements) SetBoxes(v []*int32) *RecognizeVehiclePartsResponseBodyDataElements {
	s.Boxes = v
	return s
}

func (s *RecognizeVehiclePartsResponseBodyDataElements) SetScore(v float32) *RecognizeVehiclePartsResponseBodyDataElements {
	s.Score = &v
	return s
}

func (s *RecognizeVehiclePartsResponseBodyDataElements) SetType(v string) *RecognizeVehiclePartsResponseBodyDataElements {
	s.Type = &v
	return s
}

type RecognizeVehiclePartsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeVehiclePartsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeVehiclePartsResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehiclePartsResponse) GoString() string {
	return s.String()
}

func (s *RecognizeVehiclePartsResponse) SetHeaders(v map[string]*string) *RecognizeVehiclePartsResponse {
	s.Headers = v
	return s
}

func (s *RecognizeVehiclePartsResponse) SetStatusCode(v int32) *RecognizeVehiclePartsResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeVehiclePartsResponse) SetBody(v *RecognizeVehiclePartsResponseBody) *RecognizeVehiclePartsResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("objectdet"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ClassifyVehicleInsuranceWithOptions(request *ClassifyVehicleInsuranceRequest, runtime *util.RuntimeOptions) (_result *ClassifyVehicleInsuranceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ClassifyVehicleInsurance"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ClassifyVehicleInsuranceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ClassifyVehicleInsurance(request *ClassifyVehicleInsuranceRequest) (_result *ClassifyVehicleInsuranceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ClassifyVehicleInsuranceResponse{}
	_body, _err := client.ClassifyVehicleInsuranceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ClassifyVehicleInsuranceAdvance(request *ClassifyVehicleInsuranceAdvanceRequest, runtime *util.RuntimeOptions) (_result *ClassifyVehicleInsuranceResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("objectdet"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	classifyVehicleInsuranceReq := &ClassifyVehicleInsuranceRequest{}
	openapiutil.Convert(request, classifyVehicleInsuranceReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		classifyVehicleInsuranceReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	classifyVehicleInsuranceResp, _err := client.ClassifyVehicleInsuranceWithOptions(classifyVehicleInsuranceReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = classifyVehicleInsuranceResp
	return _result, _err
}

func (client *Client) DetectIPCObjectWithOptions(request *DetectIPCObjectRequest, runtime *util.RuntimeOptions) (_result *DetectIPCObjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectIPCObject"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectIPCObjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectIPCObject(request *DetectIPCObjectRequest) (_result *DetectIPCObjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectIPCObjectResponse{}
	_body, _err := client.DetectIPCObjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectIPCObjectAdvance(request *DetectIPCObjectAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectIPCObjectResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("objectdet"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	detectIPCObjectReq := &DetectIPCObjectRequest{}
	openapiutil.Convert(request, detectIPCObjectReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		detectIPCObjectReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	detectIPCObjectResp, _err := client.DetectIPCObjectWithOptions(detectIPCObjectReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectIPCObjectResp
	return _result, _err
}

func (client *Client) DetectKitchenAnimalsWithOptions(request *DetectKitchenAnimalsRequest, runtime *util.RuntimeOptions) (_result *DetectKitchenAnimalsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURLA)) {
		body["ImageURLA"] = request.ImageURLA
	}

	if !tea.BoolValue(util.IsUnset(request.ImageURLB)) {
		body["ImageURLB"] = request.ImageURLB
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectKitchenAnimals"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectKitchenAnimalsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectKitchenAnimals(request *DetectKitchenAnimalsRequest) (_result *DetectKitchenAnimalsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectKitchenAnimalsResponse{}
	_body, _err := client.DetectKitchenAnimalsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectKitchenAnimalsAdvance(request *DetectKitchenAnimalsAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectKitchenAnimalsResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("objectdet"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	detectKitchenAnimalsReq := &DetectKitchenAnimalsRequest{}
	openapiutil.Convert(request, detectKitchenAnimalsReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLAObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLAObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		detectKitchenAnimalsReq.ImageURLA = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	if !tea.BoolValue(util.IsUnset(request.ImageURLBObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLBObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		detectKitchenAnimalsReq.ImageURLB = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	detectKitchenAnimalsResp, _err := client.DetectKitchenAnimalsWithOptions(detectKitchenAnimalsReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectKitchenAnimalsResp
	return _result, _err
}

func (client *Client) DetectMainBodyWithOptions(request *DetectMainBodyRequest, runtime *util.RuntimeOptions) (_result *DetectMainBodyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		query["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectMainBody"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectMainBodyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectMainBody(request *DetectMainBodyRequest) (_result *DetectMainBodyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectMainBodyResponse{}
	_body, _err := client.DetectMainBodyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectMainBodyAdvance(request *DetectMainBodyAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectMainBodyResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("objectdet"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	detectMainBodyReq := &DetectMainBodyRequest{}
	openapiutil.Convert(request, detectMainBodyReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		detectMainBodyReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	detectMainBodyResp, _err := client.DetectMainBodyWithOptions(detectMainBodyReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectMainBodyResp
	return _result, _err
}

func (client *Client) DetectObjectWithOptions(request *DetectObjectRequest, runtime *util.RuntimeOptions) (_result *DetectObjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectObject"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectObjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectObject(request *DetectObjectRequest) (_result *DetectObjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectObjectResponse{}
	_body, _err := client.DetectObjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectObjectAdvance(request *DetectObjectAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectObjectResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("objectdet"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	detectObjectReq := &DetectObjectRequest{}
	openapiutil.Convert(request, detectObjectReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		detectObjectReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	detectObjectResp, _err := client.DetectObjectWithOptions(detectObjectReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectObjectResp
	return _result, _err
}

func (client *Client) DetectTransparentImageWithOptions(request *DetectTransparentImageRequest, runtime *util.RuntimeOptions) (_result *DetectTransparentImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectTransparentImage"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectTransparentImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectTransparentImage(request *DetectTransparentImageRequest) (_result *DetectTransparentImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectTransparentImageResponse{}
	_body, _err := client.DetectTransparentImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectTransparentImageAdvance(request *DetectTransparentImageAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectTransparentImageResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("objectdet"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	detectTransparentImageReq := &DetectTransparentImageRequest{}
	openapiutil.Convert(request, detectTransparentImageReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		detectTransparentImageReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	detectTransparentImageResp, _err := client.DetectTransparentImageWithOptions(detectTransparentImageReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectTransparentImageResp
	return _result, _err
}

func (client *Client) DetectVehicleWithOptions(request *DetectVehicleRequest, runtime *util.RuntimeOptions) (_result *DetectVehicleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectVehicle"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectVehicleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectVehicle(request *DetectVehicleRequest) (_result *DetectVehicleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectVehicleResponse{}
	_body, _err := client.DetectVehicleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectVehicleAdvance(request *DetectVehicleAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectVehicleResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("objectdet"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	detectVehicleReq := &DetectVehicleRequest{}
	openapiutil.Convert(request, detectVehicleReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		detectVehicleReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	detectVehicleResp, _err := client.DetectVehicleWithOptions(detectVehicleReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectVehicleResp
	return _result, _err
}

func (client *Client) DetectVehicleICongestionWithOptions(tmpReq *DetectVehicleICongestionRequest, runtime *util.RuntimeOptions) (_result *DetectVehicleICongestionResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DetectVehicleICongestionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.PreRegionIntersectFeatures)) {
		request.PreRegionIntersectFeaturesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PreRegionIntersectFeatures, tea.String("PreRegionIntersectFeatures"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.RoadRegions)) {
		request.RoadRegionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RoadRegions, tea.String("RoadRegions"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	if !tea.BoolValue(util.IsUnset(request.PreRegionIntersectFeaturesShrink)) {
		body["PreRegionIntersectFeatures"] = request.PreRegionIntersectFeaturesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RoadRegionsShrink)) {
		body["RoadRegions"] = request.RoadRegionsShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectVehicleICongestion"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectVehicleICongestionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectVehicleICongestion(request *DetectVehicleICongestionRequest) (_result *DetectVehicleICongestionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectVehicleICongestionResponse{}
	_body, _err := client.DetectVehicleICongestionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectVehicleICongestionAdvance(request *DetectVehicleICongestionAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectVehicleICongestionResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("objectdet"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	detectVehicleICongestionReq := &DetectVehicleICongestionRequest{}
	openapiutil.Convert(request, detectVehicleICongestionReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		detectVehicleICongestionReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	detectVehicleICongestionResp, _err := client.DetectVehicleICongestionWithOptions(detectVehicleICongestionReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectVehicleICongestionResp
	return _result, _err
}

func (client *Client) DetectVehicleIllegalParkingWithOptions(tmpReq *DetectVehicleIllegalParkingRequest, runtime *util.RuntimeOptions) (_result *DetectVehicleIllegalParkingResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DetectVehicleIllegalParkingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RoadRegions)) {
		request.RoadRegionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RoadRegions, tea.String("RoadRegions"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	if !tea.BoolValue(util.IsUnset(request.RoadRegionsShrink)) {
		body["RoadRegions"] = request.RoadRegionsShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectVehicleIllegalParking"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectVehicleIllegalParkingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectVehicleIllegalParking(request *DetectVehicleIllegalParkingRequest) (_result *DetectVehicleIllegalParkingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectVehicleIllegalParkingResponse{}
	_body, _err := client.DetectVehicleIllegalParkingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectVehicleIllegalParkingAdvance(request *DetectVehicleIllegalParkingAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectVehicleIllegalParkingResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("objectdet"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	detectVehicleIllegalParkingReq := &DetectVehicleIllegalParkingRequest{}
	openapiutil.Convert(request, detectVehicleIllegalParkingReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		detectVehicleIllegalParkingReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	detectVehicleIllegalParkingResp, _err := client.DetectVehicleIllegalParkingWithOptions(detectVehicleIllegalParkingReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectVehicleIllegalParkingResp
	return _result, _err
}

func (client *Client) DetectVideoIPCObjectWithOptions(request *DetectVideoIPCObjectRequest, runtime *util.RuntimeOptions) (_result *DetectVideoIPCObjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallbackOnlyHasObject)) {
		body["CallbackOnlyHasObject"] = request.CallbackOnlyHasObject
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimestamp)) {
		body["StartTimestamp"] = request.StartTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.VideoURL)) {
		body["VideoURL"] = request.VideoURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectVideoIPCObject"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectVideoIPCObjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectVideoIPCObject(request *DetectVideoIPCObjectRequest) (_result *DetectVideoIPCObjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectVideoIPCObjectResponse{}
	_body, _err := client.DetectVideoIPCObjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectVideoIPCObjectAdvance(request *DetectVideoIPCObjectAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectVideoIPCObjectResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("objectdet"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	detectVideoIPCObjectReq := &DetectVideoIPCObjectRequest{}
	openapiutil.Convert(request, detectVideoIPCObjectReq)
	if !tea.BoolValue(util.IsUnset(request.VideoURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.VideoURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		detectVideoIPCObjectReq.VideoURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	detectVideoIPCObjectResp, _err := client.DetectVideoIPCObjectWithOptions(detectVideoIPCObjectReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectVideoIPCObjectResp
	return _result, _err
}

func (client *Client) DetectWhiteBaseImageWithOptions(request *DetectWhiteBaseImageRequest, runtime *util.RuntimeOptions) (_result *DetectWhiteBaseImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectWhiteBaseImage"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectWhiteBaseImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectWhiteBaseImage(request *DetectWhiteBaseImageRequest) (_result *DetectWhiteBaseImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectWhiteBaseImageResponse{}
	_body, _err := client.DetectWhiteBaseImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectWhiteBaseImageAdvance(request *DetectWhiteBaseImageAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectWhiteBaseImageResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("objectdet"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	detectWhiteBaseImageReq := &DetectWhiteBaseImageRequest{}
	openapiutil.Convert(request, detectWhiteBaseImageReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		detectWhiteBaseImageReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	detectWhiteBaseImageResp, _err := client.DetectWhiteBaseImageWithOptions(detectWhiteBaseImageReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectWhiteBaseImageResp
	return _result, _err
}

func (client *Client) DetectWorkwearWithOptions(tmpReq *DetectWorkwearRequest, runtime *util.RuntimeOptions) (_result *DetectWorkwearResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DetectWorkwearShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Clothes)) {
		request.ClothesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Clothes, tea.String("Clothes"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClothesShrink)) {
		body["Clothes"] = request.ClothesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["ImageUrl"] = request.ImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		body["Labels"] = request.Labels
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectWorkwear"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectWorkwearResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectWorkwear(request *DetectWorkwearRequest) (_result *DetectWorkwearResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectWorkwearResponse{}
	_body, _err := client.DetectWorkwearWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectWorkwearAdvance(request *DetectWorkwearAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectWorkwearResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("objectdet"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	detectWorkwearReq := &DetectWorkwearRequest{}
	openapiutil.Convert(request, detectWorkwearReq)
	if !tea.BoolValue(util.IsUnset(request.ImageUrlObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageUrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		detectWorkwearReq.ImageUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	detectWorkwearResp, _err := client.DetectWorkwearWithOptions(detectWorkwearReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectWorkwearResp
	return _result, _err
}

func (client *Client) GetAsyncJobResultWithOptions(request *GetAsyncJobResultRequest, runtime *util.RuntimeOptions) (_result *GetAsyncJobResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAsyncJobResult"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAsyncJobResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAsyncJobResult(request *GetAsyncJobResultRequest) (_result *GetAsyncJobResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAsyncJobResultResponse{}
	_body, _err := client.GetAsyncJobResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeVehicleDamageWithOptions(request *RecognizeVehicleDamageRequest, runtime *util.RuntimeOptions) (_result *RecognizeVehicleDamageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeVehicleDamage"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeVehicleDamageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeVehicleDamage(request *RecognizeVehicleDamageRequest) (_result *RecognizeVehicleDamageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeVehicleDamageResponse{}
	_body, _err := client.RecognizeVehicleDamageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeVehicleDamageAdvance(request *RecognizeVehicleDamageAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeVehicleDamageResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("objectdet"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	recognizeVehicleDamageReq := &RecognizeVehicleDamageRequest{}
	openapiutil.Convert(request, recognizeVehicleDamageReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizeVehicleDamageReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recognizeVehicleDamageResp, _err := client.RecognizeVehicleDamageWithOptions(recognizeVehicleDamageReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeVehicleDamageResp
	return _result, _err
}

func (client *Client) RecognizeVehicleDashboardWithOptions(request *RecognizeVehicleDashboardRequest, runtime *util.RuntimeOptions) (_result *RecognizeVehicleDashboardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeVehicleDashboard"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeVehicleDashboardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeVehicleDashboard(request *RecognizeVehicleDashboardRequest) (_result *RecognizeVehicleDashboardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeVehicleDashboardResponse{}
	_body, _err := client.RecognizeVehicleDashboardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeVehicleDashboardAdvance(request *RecognizeVehicleDashboardAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeVehicleDashboardResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("objectdet"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	recognizeVehicleDashboardReq := &RecognizeVehicleDashboardRequest{}
	openapiutil.Convert(request, recognizeVehicleDashboardReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizeVehicleDashboardReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recognizeVehicleDashboardResp, _err := client.RecognizeVehicleDashboardWithOptions(recognizeVehicleDashboardReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeVehicleDashboardResp
	return _result, _err
}

func (client *Client) RecognizeVehiclePartsWithOptions(request *RecognizeVehiclePartsRequest, runtime *util.RuntimeOptions) (_result *RecognizeVehiclePartsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeVehicleParts"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeVehiclePartsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeVehicleParts(request *RecognizeVehiclePartsRequest) (_result *RecognizeVehiclePartsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeVehiclePartsResponse{}
	_body, _err := client.RecognizeVehiclePartsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeVehiclePartsAdvance(request *RecognizeVehiclePartsAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeVehiclePartsResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("objectdet"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	recognizeVehiclePartsReq := &RecognizeVehiclePartsRequest{}
	openapiutil.Convert(request, recognizeVehiclePartsReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizeVehiclePartsReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recognizeVehiclePartsResp, _err := client.RecognizeVehiclePartsWithOptions(recognizeVehiclePartsReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeVehiclePartsResp
	return _result, _err
}
