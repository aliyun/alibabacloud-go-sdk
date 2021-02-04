// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	openplatform "github.com/alibabacloud-go/openplatform-20191219/client"
	fileform "github.com/alibabacloud-go/tea-fileform/service"
	oss "github.com/alibabacloud-go/tea-oss-sdk/client"
	ossutil "github.com/alibabacloud-go/tea-oss-utils/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"io"
)

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
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
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
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DetectTransparentImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DetectTransparentImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectTransparentImageResponseBody) GoString() string {
	return s.String()
}

func (s *DetectTransparentImageResponseBody) SetRequestId(v string) *DetectTransparentImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectTransparentImageResponseBody) SetData(v *DetectTransparentImageResponseBodyData) *DetectTransparentImageResponseBody {
	s.Data = v
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
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetectTransparentImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DetectTransparentImageResponse) SetBody(v *DetectTransparentImageResponseBody) *DetectTransparentImageResponse {
	s.Body = v
	return s
}

type DetectVehicleICongestionRequest struct {
	// A short description of struct
	ImageURL                   *string                                                      `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	RoadRegions                []*DetectVehicleICongestionRequestRoadRegions                `json:"RoadRegions,omitempty" xml:"RoadRegions,omitempty" type:"Repeated"`
	PreRegionIntersectFeatures []*DetectVehicleICongestionRequestPreRegionIntersectFeatures `json:"PreRegionIntersectFeatures,omitempty" xml:"PreRegionIntersectFeatures,omitempty" type:"Repeated"`
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

func (s *DetectVehicleICongestionRequest) SetRoadRegions(v []*DetectVehicleICongestionRequestRoadRegions) *DetectVehicleICongestionRequest {
	s.RoadRegions = v
	return s
}

func (s *DetectVehicleICongestionRequest) SetPreRegionIntersectFeatures(v []*DetectVehicleICongestionRequestPreRegionIntersectFeatures) *DetectVehicleICongestionRequest {
	s.PreRegionIntersectFeatures = v
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

type DetectVehicleICongestionShrinkRequest struct {
	// A short description of struct
	ImageURL                         *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	RoadRegionsShrink                *string `json:"RoadRegions,omitempty" xml:"RoadRegions,omitempty"`
	PreRegionIntersectFeaturesShrink *string `json:"PreRegionIntersectFeatures,omitempty" xml:"PreRegionIntersectFeatures,omitempty"`
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

func (s *DetectVehicleICongestionShrinkRequest) SetRoadRegionsShrink(v string) *DetectVehicleICongestionShrinkRequest {
	s.RoadRegionsShrink = &v
	return s
}

func (s *DetectVehicleICongestionShrinkRequest) SetPreRegionIntersectFeaturesShrink(v string) *DetectVehicleICongestionShrinkRequest {
	s.PreRegionIntersectFeaturesShrink = &v
	return s
}

type DetectVehicleICongestionResponseBody struct {
	// Id of the request
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DetectVehicleICongestionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DetectVehicleICongestionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleICongestionResponseBody) GoString() string {
	return s.String()
}

func (s *DetectVehicleICongestionResponseBody) SetRequestId(v string) *DetectVehicleICongestionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectVehicleICongestionResponseBody) SetData(v *DetectVehicleICongestionResponseBodyData) *DetectVehicleICongestionResponseBody {
	s.Data = v
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

func (s *DetectVehicleICongestionResponseBodyDataElements) SetScore(v float32) *DetectVehicleICongestionResponseBodyDataElements {
	s.Score = &v
	return s
}

func (s *DetectVehicleICongestionResponseBodyDataElements) SetTypeName(v string) *DetectVehicleICongestionResponseBodyDataElements {
	s.TypeName = &v
	return s
}

type DetectVehicleICongestionResponseBodyDataElementsBoxes struct {
	Left   *int64 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int64 `json:"Top,omitempty" xml:"Top,omitempty"`
	Right  *int64 `json:"Right,omitempty" xml:"Right,omitempty"`
	Bottom *int64 `json:"Bottom,omitempty" xml:"Bottom,omitempty"`
}

func (s DetectVehicleICongestionResponseBodyDataElementsBoxes) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleICongestionResponseBodyDataElementsBoxes) GoString() string {
	return s.String()
}

func (s *DetectVehicleICongestionResponseBodyDataElementsBoxes) SetLeft(v int64) *DetectVehicleICongestionResponseBodyDataElementsBoxes {
	s.Left = &v
	return s
}

func (s *DetectVehicleICongestionResponseBodyDataElementsBoxes) SetTop(v int64) *DetectVehicleICongestionResponseBodyDataElementsBoxes {
	s.Top = &v
	return s
}

func (s *DetectVehicleICongestionResponseBodyDataElementsBoxes) SetRight(v int64) *DetectVehicleICongestionResponseBodyDataElementsBoxes {
	s.Right = &v
	return s
}

func (s *DetectVehicleICongestionResponseBodyDataElementsBoxes) SetBottom(v int64) *DetectVehicleICongestionResponseBodyDataElementsBoxes {
	s.Bottom = &v
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
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetectVehicleICongestionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DetectVehicleICongestionResponse) SetBody(v *DetectVehicleICongestionResponseBody) *DetectVehicleICongestionResponse {
	s.Body = v
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
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
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
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ClassifyVehicleInsuranceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ClassifyVehicleInsuranceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ClassifyVehicleInsuranceResponseBody) GoString() string {
	return s.String()
}

func (s *ClassifyVehicleInsuranceResponseBody) SetRequestId(v string) *ClassifyVehicleInsuranceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClassifyVehicleInsuranceResponseBody) SetData(v *ClassifyVehicleInsuranceResponseBodyData) *ClassifyVehicleInsuranceResponseBody {
	s.Data = v
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
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	Name  *string  `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ClassifyVehicleInsuranceResponseBodyDataLabels) String() string {
	return tea.Prettify(s)
}

func (s ClassifyVehicleInsuranceResponseBodyDataLabels) GoString() string {
	return s.String()
}

func (s *ClassifyVehicleInsuranceResponseBodyDataLabels) SetScore(v float32) *ClassifyVehicleInsuranceResponseBodyDataLabels {
	s.Score = &v
	return s
}

func (s *ClassifyVehicleInsuranceResponseBodyDataLabels) SetName(v string) *ClassifyVehicleInsuranceResponseBodyDataLabels {
	s.Name = &v
	return s
}

type ClassifyVehicleInsuranceResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ClassifyVehicleInsuranceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ClassifyVehicleInsuranceResponse) SetBody(v *ClassifyVehicleInsuranceResponseBody) *ClassifyVehicleInsuranceResponse {
	s.Body = v
	return s
}

type DetectIPCObjectRequest struct {
	// 图片URL地址
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

type DetectIPCObjectResponseBody struct {
	// Id of the request
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DetectIPCObjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DetectIPCObjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectIPCObjectResponseBody) GoString() string {
	return s.String()
}

func (s *DetectIPCObjectResponseBody) SetRequestId(v string) *DetectIPCObjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectIPCObjectResponseBody) SetData(v *DetectIPCObjectResponseBodyData) *DetectIPCObjectResponseBody {
	s.Data = v
	return s
}

type DetectIPCObjectResponseBodyData struct {
	Elements []*DetectIPCObjectResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	Width    *int64                                     `json:"Width,omitempty" xml:"Width,omitempty"`
	Height   *int64                                     `json:"Height,omitempty" xml:"Height,omitempty"`
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

func (s *DetectIPCObjectResponseBodyData) SetWidth(v int64) *DetectIPCObjectResponseBodyData {
	s.Width = &v
	return s
}

func (s *DetectIPCObjectResponseBodyData) SetHeight(v int64) *DetectIPCObjectResponseBodyData {
	s.Height = &v
	return s
}

type DetectIPCObjectResponseBodyDataElements struct {
	Type  *string  `json:"Type,omitempty" xml:"Type,omitempty"`
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	Box   []*int64 `json:"Box,omitempty" xml:"Box,omitempty" type:"Repeated"`
}

func (s DetectIPCObjectResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s DetectIPCObjectResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *DetectIPCObjectResponseBodyDataElements) SetType(v string) *DetectIPCObjectResponseBodyDataElements {
	s.Type = &v
	return s
}

func (s *DetectIPCObjectResponseBodyDataElements) SetScore(v float32) *DetectIPCObjectResponseBodyDataElements {
	s.Score = &v
	return s
}

func (s *DetectIPCObjectResponseBodyDataElements) SetBox(v []*int64) *DetectIPCObjectResponseBodyDataElements {
	s.Box = v
	return s
}

type DetectIPCObjectResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetectIPCObjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DetectIPCObjectResponse) SetBody(v *DetectIPCObjectResponseBody) *DetectIPCObjectResponse {
	s.Body = v
	return s
}

type GetVehicleRepairPlanRequest struct {
	TaskId         *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	CarNumberImage *string `json:"CarNumberImage,omitempty" xml:"CarNumberImage,omitempty"`
	VinCodeImage   *string `json:"VinCodeImage,omitempty" xml:"VinCodeImage,omitempty"`
}

func (s GetVehicleRepairPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVehicleRepairPlanRequest) GoString() string {
	return s.String()
}

func (s *GetVehicleRepairPlanRequest) SetTaskId(v string) *GetVehicleRepairPlanRequest {
	s.TaskId = &v
	return s
}

func (s *GetVehicleRepairPlanRequest) SetCarNumberImage(v string) *GetVehicleRepairPlanRequest {
	s.CarNumberImage = &v
	return s
}

func (s *GetVehicleRepairPlanRequest) SetVinCodeImage(v string) *GetVehicleRepairPlanRequest {
	s.VinCodeImage = &v
	return s
}

type GetVehicleRepairPlanResponseBody struct {
	HttpCode     *int32                                `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	RequestId    *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *GetVehicleRepairPlanResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                               `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetVehicleRepairPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVehicleRepairPlanResponseBody) GoString() string {
	return s.String()
}

func (s *GetVehicleRepairPlanResponseBody) SetHttpCode(v int32) *GetVehicleRepairPlanResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetVehicleRepairPlanResponseBody) SetRequestId(v string) *GetVehicleRepairPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVehicleRepairPlanResponseBody) SetData(v *GetVehicleRepairPlanResponseBodyData) *GetVehicleRepairPlanResponseBody {
	s.Data = v
	return s
}

func (s *GetVehicleRepairPlanResponseBody) SetErrorMessage(v string) *GetVehicleRepairPlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetVehicleRepairPlanResponseBody) SetCode(v string) *GetVehicleRepairPlanResponseBody {
	s.Code = &v
	return s
}

func (s *GetVehicleRepairPlanResponseBody) SetSuccess(v bool) *GetVehicleRepairPlanResponseBody {
	s.Success = &v
	return s
}

type GetVehicleRepairPlanResponseBodyData struct {
	RepairParts []*GetVehicleRepairPlanResponseBodyDataRepairParts `json:"RepairParts,omitempty" xml:"RepairParts,omitempty" type:"Repeated"`
	FrameNo     *string                                            `json:"FrameNo,omitempty" xml:"FrameNo,omitempty"`
}

func (s GetVehicleRepairPlanResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetVehicleRepairPlanResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetVehicleRepairPlanResponseBodyData) SetRepairParts(v []*GetVehicleRepairPlanResponseBodyDataRepairParts) *GetVehicleRepairPlanResponseBodyData {
	s.RepairParts = v
	return s
}

func (s *GetVehicleRepairPlanResponseBodyData) SetFrameNo(v string) *GetVehicleRepairPlanResponseBodyData {
	s.FrameNo = &v
	return s
}

type GetVehicleRepairPlanResponseBodyDataRepairParts struct {
	RelationType         *string `json:"RelationType,omitempty" xml:"RelationType,omitempty"`
	PartsStdCode         *string `json:"PartsStdCode,omitempty" xml:"PartsStdCode,omitempty"`
	PartNameMatch        *bool   `json:"PartNameMatch,omitempty" xml:"PartNameMatch,omitempty"`
	RepairFee            *string `json:"RepairFee,omitempty" xml:"RepairFee,omitempty"`
	OutStandardPartsName *string `json:"OutStandardPartsName,omitempty" xml:"OutStandardPartsName,omitempty"`
	PartsStdName         *string `json:"PartsStdName,omitempty" xml:"PartsStdName,omitempty"`
	RepairTypeName       *string `json:"RepairTypeName,omitempty" xml:"RepairTypeName,omitempty"`
	RepairType           *string `json:"RepairType,omitempty" xml:"RepairType,omitempty"`
	OeMatch              *bool   `json:"OeMatch,omitempty" xml:"OeMatch,omitempty"`
	OutStandardPartsId   *string `json:"OutStandardPartsId,omitempty" xml:"OutStandardPartsId,omitempty"`
	GarageType           *string `json:"GarageType,omitempty" xml:"GarageType,omitempty"`
}

func (s GetVehicleRepairPlanResponseBodyDataRepairParts) String() string {
	return tea.Prettify(s)
}

func (s GetVehicleRepairPlanResponseBodyDataRepairParts) GoString() string {
	return s.String()
}

func (s *GetVehicleRepairPlanResponseBodyDataRepairParts) SetRelationType(v string) *GetVehicleRepairPlanResponseBodyDataRepairParts {
	s.RelationType = &v
	return s
}

func (s *GetVehicleRepairPlanResponseBodyDataRepairParts) SetPartsStdCode(v string) *GetVehicleRepairPlanResponseBodyDataRepairParts {
	s.PartsStdCode = &v
	return s
}

func (s *GetVehicleRepairPlanResponseBodyDataRepairParts) SetPartNameMatch(v bool) *GetVehicleRepairPlanResponseBodyDataRepairParts {
	s.PartNameMatch = &v
	return s
}

func (s *GetVehicleRepairPlanResponseBodyDataRepairParts) SetRepairFee(v string) *GetVehicleRepairPlanResponseBodyDataRepairParts {
	s.RepairFee = &v
	return s
}

func (s *GetVehicleRepairPlanResponseBodyDataRepairParts) SetOutStandardPartsName(v string) *GetVehicleRepairPlanResponseBodyDataRepairParts {
	s.OutStandardPartsName = &v
	return s
}

func (s *GetVehicleRepairPlanResponseBodyDataRepairParts) SetPartsStdName(v string) *GetVehicleRepairPlanResponseBodyDataRepairParts {
	s.PartsStdName = &v
	return s
}

func (s *GetVehicleRepairPlanResponseBodyDataRepairParts) SetRepairTypeName(v string) *GetVehicleRepairPlanResponseBodyDataRepairParts {
	s.RepairTypeName = &v
	return s
}

func (s *GetVehicleRepairPlanResponseBodyDataRepairParts) SetRepairType(v string) *GetVehicleRepairPlanResponseBodyDataRepairParts {
	s.RepairType = &v
	return s
}

func (s *GetVehicleRepairPlanResponseBodyDataRepairParts) SetOeMatch(v bool) *GetVehicleRepairPlanResponseBodyDataRepairParts {
	s.OeMatch = &v
	return s
}

func (s *GetVehicleRepairPlanResponseBodyDataRepairParts) SetOutStandardPartsId(v string) *GetVehicleRepairPlanResponseBodyDataRepairParts {
	s.OutStandardPartsId = &v
	return s
}

func (s *GetVehicleRepairPlanResponseBodyDataRepairParts) SetGarageType(v string) *GetVehicleRepairPlanResponseBodyDataRepairParts {
	s.GarageType = &v
	return s
}

type GetVehicleRepairPlanResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetVehicleRepairPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetVehicleRepairPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVehicleRepairPlanResponse) GoString() string {
	return s.String()
}

func (s *GetVehicleRepairPlanResponse) SetHeaders(v map[string]*string) *GetVehicleRepairPlanResponse {
	s.Headers = v
	return s
}

func (s *GetVehicleRepairPlanResponse) SetBody(v *GetVehicleRepairPlanResponseBody) *GetVehicleRepairPlanResponse {
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
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
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
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DetectWhiteBaseImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DetectWhiteBaseImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectWhiteBaseImageResponseBody) GoString() string {
	return s.String()
}

func (s *DetectWhiteBaseImageResponseBody) SetRequestId(v string) *DetectWhiteBaseImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectWhiteBaseImageResponseBody) SetData(v *DetectWhiteBaseImageResponseBodyData) *DetectWhiteBaseImageResponseBody {
	s.Data = v
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
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetectWhiteBaseImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DetectWhiteBaseImageResponse) SetBody(v *DetectWhiteBaseImageResponseBody) *DetectWhiteBaseImageResponse {
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
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
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
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DetectMainBodyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DetectMainBodyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectMainBodyResponseBody) GoString() string {
	return s.String()
}

func (s *DetectMainBodyResponseBody) SetRequestId(v string) *DetectMainBodyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectMainBodyResponseBody) SetData(v *DetectMainBodyResponseBodyData) *DetectMainBodyResponseBody {
	s.Data = v
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
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Y      *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
	X      *int32 `json:"X,omitempty" xml:"X,omitempty"`
}

func (s DetectMainBodyResponseBodyDataLocation) String() string {
	return tea.Prettify(s)
}

func (s DetectMainBodyResponseBodyDataLocation) GoString() string {
	return s.String()
}

func (s *DetectMainBodyResponseBodyDataLocation) SetWidth(v int32) *DetectMainBodyResponseBodyDataLocation {
	s.Width = &v
	return s
}

func (s *DetectMainBodyResponseBodyDataLocation) SetHeight(v int32) *DetectMainBodyResponseBodyDataLocation {
	s.Height = &v
	return s
}

func (s *DetectMainBodyResponseBodyDataLocation) SetY(v int32) *DetectMainBodyResponseBodyDataLocation {
	s.Y = &v
	return s
}

func (s *DetectMainBodyResponseBodyDataLocation) SetX(v int32) *DetectMainBodyResponseBodyDataLocation {
	s.X = &v
	return s
}

type DetectMainBodyResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetectMainBodyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DetectMainBodyResponse) SetBody(v *DetectMainBodyResponseBody) *DetectMainBodyResponse {
	s.Body = v
	return s
}

type DetectVehicleRequest struct {
	ImageType *int32  `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	ImageURL  *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectVehicleRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleRequest) GoString() string {
	return s.String()
}

func (s *DetectVehicleRequest) SetImageType(v int32) *DetectVehicleRequest {
	s.ImageType = &v
	return s
}

func (s *DetectVehicleRequest) SetImageURL(v string) *DetectVehicleRequest {
	s.ImageURL = &v
	return s
}

type DetectVehicleAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
	ImageType      *int32    `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
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

func (s *DetectVehicleAdvanceRequest) SetImageType(v int32) *DetectVehicleAdvanceRequest {
	s.ImageType = &v
	return s
}

type DetectVehicleResponseBody struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DetectVehicleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DetectVehicleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleResponseBody) GoString() string {
	return s.String()
}

func (s *DetectVehicleResponseBody) SetRequestId(v string) *DetectVehicleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectVehicleResponseBody) SetData(v *DetectVehicleResponseBodyData) *DetectVehicleResponseBody {
	s.Data = v
	return s
}

type DetectVehicleResponseBodyData struct {
	DetectObjectInfoList []*DetectVehicleResponseBodyDataDetectObjectInfoList `json:"DetectObjectInfoList,omitempty" xml:"DetectObjectInfoList,omitempty" type:"Repeated"`
	Width                *int32                                               `json:"Width,omitempty" xml:"Width,omitempty"`
	Height               *int32                                               `json:"Height,omitempty" xml:"Height,omitempty"`
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

func (s *DetectVehicleResponseBodyData) SetWidth(v int32) *DetectVehicleResponseBodyData {
	s.Width = &v
	return s
}

func (s *DetectVehicleResponseBodyData) SetHeight(v int32) *DetectVehicleResponseBodyData {
	s.Height = &v
	return s
}

type DetectVehicleResponseBodyDataDetectObjectInfoList struct {
	Type  *string  `json:"Type,omitempty" xml:"Type,omitempty"`
	Boxes []*int32 `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	Id    *int32   `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DetectVehicleResponseBodyDataDetectObjectInfoList) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleResponseBodyDataDetectObjectInfoList) GoString() string {
	return s.String()
}

func (s *DetectVehicleResponseBodyDataDetectObjectInfoList) SetType(v string) *DetectVehicleResponseBodyDataDetectObjectInfoList {
	s.Type = &v
	return s
}

func (s *DetectVehicleResponseBodyDataDetectObjectInfoList) SetBoxes(v []*int32) *DetectVehicleResponseBodyDataDetectObjectInfoList {
	s.Boxes = v
	return s
}

func (s *DetectVehicleResponseBodyDataDetectObjectInfoList) SetScore(v float32) *DetectVehicleResponseBodyDataDetectObjectInfoList {
	s.Score = &v
	return s
}

func (s *DetectVehicleResponseBodyDataDetectObjectInfoList) SetId(v int32) *DetectVehicleResponseBodyDataDetectObjectInfoList {
	s.Id = &v
	return s
}

type DetectVehicleResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetectVehicleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DetectVehicleResponse) SetBody(v *DetectVehicleResponseBody) *DetectVehicleResponse {
	s.Body = v
	return s
}

type DetectVehicleIllegalParkingRequest struct {
	// A short description of struct
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

type DetectVehicleIllegalParkingShrinkRequest struct {
	// A short description of struct
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
	// Id of the request
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DetectVehicleIllegalParkingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DetectVehicleIllegalParkingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleIllegalParkingResponseBody) GoString() string {
	return s.String()
}

func (s *DetectVehicleIllegalParkingResponseBody) SetRequestId(v string) *DetectVehicleIllegalParkingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectVehicleIllegalParkingResponseBody) SetData(v *DetectVehicleIllegalParkingResponseBodyData) *DetectVehicleIllegalParkingResponseBody {
	s.Data = v
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

func (s *DetectVehicleIllegalParkingResponseBodyDataElements) SetScore(v float32) *DetectVehicleIllegalParkingResponseBodyDataElements {
	s.Score = &v
	return s
}

func (s *DetectVehicleIllegalParkingResponseBodyDataElements) SetTypeName(v string) *DetectVehicleIllegalParkingResponseBodyDataElements {
	s.TypeName = &v
	return s
}

type DetectVehicleIllegalParkingResponseBodyDataElementsBoxes struct {
	Left   *int64 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int64 `json:"Top,omitempty" xml:"Top,omitempty"`
	Right  *int64 `json:"Right,omitempty" xml:"Right,omitempty"`
	Bottom *int64 `json:"Bottom,omitempty" xml:"Bottom,omitempty"`
}

func (s DetectVehicleIllegalParkingResponseBodyDataElementsBoxes) String() string {
	return tea.Prettify(s)
}

func (s DetectVehicleIllegalParkingResponseBodyDataElementsBoxes) GoString() string {
	return s.String()
}

func (s *DetectVehicleIllegalParkingResponseBodyDataElementsBoxes) SetLeft(v int64) *DetectVehicleIllegalParkingResponseBodyDataElementsBoxes {
	s.Left = &v
	return s
}

func (s *DetectVehicleIllegalParkingResponseBodyDataElementsBoxes) SetTop(v int64) *DetectVehicleIllegalParkingResponseBodyDataElementsBoxes {
	s.Top = &v
	return s
}

func (s *DetectVehicleIllegalParkingResponseBodyDataElementsBoxes) SetRight(v int64) *DetectVehicleIllegalParkingResponseBodyDataElementsBoxes {
	s.Right = &v
	return s
}

func (s *DetectVehicleIllegalParkingResponseBodyDataElementsBoxes) SetBottom(v int64) *DetectVehicleIllegalParkingResponseBodyDataElementsBoxes {
	s.Bottom = &v
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
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetectVehicleIllegalParkingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DetectVehicleIllegalParkingResponse) SetBody(v *DetectVehicleIllegalParkingResponseBody) *DetectVehicleIllegalParkingResponse {
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
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
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
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *RecognizeVehicleDamageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s RecognizeVehicleDamageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehicleDamageResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeVehicleDamageResponseBody) SetRequestId(v string) *RecognizeVehicleDamageResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeVehicleDamageResponseBody) SetData(v *RecognizeVehicleDamageResponseBodyData) *RecognizeVehicleDamageResponseBody {
	s.Data = v
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
	Type   *string    `json:"Type,omitempty" xml:"Type,omitempty"`
	Scores []*float32 `json:"Scores,omitempty" xml:"Scores,omitempty" type:"Repeated"`
	Boxes  []*int32   `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
	Score  *float32   `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s RecognizeVehicleDamageResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehicleDamageResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *RecognizeVehicleDamageResponseBodyDataElements) SetType(v string) *RecognizeVehicleDamageResponseBodyDataElements {
	s.Type = &v
	return s
}

func (s *RecognizeVehicleDamageResponseBodyDataElements) SetScores(v []*float32) *RecognizeVehicleDamageResponseBodyDataElements {
	s.Scores = v
	return s
}

func (s *RecognizeVehicleDamageResponseBodyDataElements) SetBoxes(v []*int32) *RecognizeVehicleDamageResponseBodyDataElements {
	s.Boxes = v
	return s
}

func (s *RecognizeVehicleDamageResponseBodyDataElements) SetScore(v float32) *RecognizeVehicleDamageResponseBodyDataElements {
	s.Score = &v
	return s
}

type RecognizeVehicleDamageResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeVehicleDamageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
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
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *RecognizeVehicleDashboardResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s RecognizeVehicleDashboardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehicleDashboardResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeVehicleDashboardResponseBody) SetRequestId(v string) *RecognizeVehicleDashboardResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeVehicleDashboardResponseBody) SetData(v *RecognizeVehicleDashboardResponseBodyData) *RecognizeVehicleDashboardResponseBody {
	s.Data = v
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
	Score     *float32   `json:"Score,omitempty" xml:"Score,omitempty"`
	Label     *string    `json:"Label,omitempty" xml:"Label,omitempty"`
	ClassName *string    `json:"ClassName,omitempty" xml:"ClassName,omitempty"`
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

func (s *RecognizeVehicleDashboardResponseBodyDataElements) SetScore(v float32) *RecognizeVehicleDashboardResponseBodyDataElements {
	s.Score = &v
	return s
}

func (s *RecognizeVehicleDashboardResponseBodyDataElements) SetLabel(v string) *RecognizeVehicleDashboardResponseBodyDataElements {
	s.Label = &v
	return s
}

func (s *RecognizeVehicleDashboardResponseBodyDataElements) SetClassName(v string) *RecognizeVehicleDashboardResponseBodyDataElements {
	s.ClassName = &v
	return s
}

type RecognizeVehicleDashboardResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeVehicleDashboardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
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
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *RecognizeVehiclePartsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s RecognizeVehiclePartsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehiclePartsResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeVehiclePartsResponseBody) SetRequestId(v string) *RecognizeVehiclePartsResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeVehiclePartsResponseBody) SetData(v *RecognizeVehiclePartsResponseBodyData) *RecognizeVehiclePartsResponseBody {
	s.Data = v
	return s
}

type RecognizeVehiclePartsResponseBodyData struct {
	Elements     []*RecognizeVehiclePartsResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	OriginShapes []*int32                                         `json:"OriginShapes,omitempty" xml:"OriginShapes,omitempty" type:"Repeated"`
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
	Type  *string  `json:"Type,omitempty" xml:"Type,omitempty"`
	Boxes []*int32 `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s RecognizeVehiclePartsResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehiclePartsResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *RecognizeVehiclePartsResponseBodyDataElements) SetType(v string) *RecognizeVehiclePartsResponseBodyDataElements {
	s.Type = &v
	return s
}

func (s *RecognizeVehiclePartsResponseBodyDataElements) SetBoxes(v []*int32) *RecognizeVehiclePartsResponseBodyDataElements {
	s.Boxes = v
	return s
}

func (s *RecognizeVehiclePartsResponseBodyDataElements) SetScore(v float32) *RecognizeVehiclePartsResponseBodyDataElements {
	s.Score = &v
	return s
}

type RecognizeVehiclePartsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeVehiclePartsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeVehiclePartsResponse) SetBody(v *RecognizeVehiclePartsResponseBody) *RecognizeVehiclePartsResponse {
	s.Body = v
	return s
}

type GenerateVehicleRepairPlanRequest struct {
	DamageImageList []*GenerateVehicleRepairPlanRequestDamageImageList `json:"DamageImageList,omitempty" xml:"DamageImageList,omitempty" type:"Repeated"`
}

func (s GenerateVehicleRepairPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateVehicleRepairPlanRequest) GoString() string {
	return s.String()
}

func (s *GenerateVehicleRepairPlanRequest) SetDamageImageList(v []*GenerateVehicleRepairPlanRequestDamageImageList) *GenerateVehicleRepairPlanRequest {
	s.DamageImageList = v
	return s
}

type GenerateVehicleRepairPlanRequestDamageImageList struct {
	CreateTimeStamp *string `json:"CreateTimeStamp,omitempty" xml:"CreateTimeStamp,omitempty"`
	ImageUrl        *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s GenerateVehicleRepairPlanRequestDamageImageList) String() string {
	return tea.Prettify(s)
}

func (s GenerateVehicleRepairPlanRequestDamageImageList) GoString() string {
	return s.String()
}

func (s *GenerateVehicleRepairPlanRequestDamageImageList) SetCreateTimeStamp(v string) *GenerateVehicleRepairPlanRequestDamageImageList {
	s.CreateTimeStamp = &v
	return s
}

func (s *GenerateVehicleRepairPlanRequestDamageImageList) SetImageUrl(v string) *GenerateVehicleRepairPlanRequestDamageImageList {
	s.ImageUrl = &v
	return s
}

type GenerateVehicleRepairPlanResponseBody struct {
	HttpCode     *int32                                     `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	RequestId    *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *GenerateVehicleRepairPlanResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                    `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GenerateVehicleRepairPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateVehicleRepairPlanResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateVehicleRepairPlanResponseBody) SetHttpCode(v int32) *GenerateVehicleRepairPlanResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GenerateVehicleRepairPlanResponseBody) SetRequestId(v string) *GenerateVehicleRepairPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateVehicleRepairPlanResponseBody) SetData(v *GenerateVehicleRepairPlanResponseBodyData) *GenerateVehicleRepairPlanResponseBody {
	s.Data = v
	return s
}

func (s *GenerateVehicleRepairPlanResponseBody) SetErrorMessage(v string) *GenerateVehicleRepairPlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GenerateVehicleRepairPlanResponseBody) SetCode(v string) *GenerateVehicleRepairPlanResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateVehicleRepairPlanResponseBody) SetSuccess(v bool) *GenerateVehicleRepairPlanResponseBody {
	s.Success = &v
	return s
}

type GenerateVehicleRepairPlanResponseBodyData struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GenerateVehicleRepairPlanResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GenerateVehicleRepairPlanResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateVehicleRepairPlanResponseBodyData) SetTaskId(v string) *GenerateVehicleRepairPlanResponseBodyData {
	s.TaskId = &v
	return s
}

type GenerateVehicleRepairPlanResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GenerateVehicleRepairPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateVehicleRepairPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateVehicleRepairPlanResponse) GoString() string {
	return s.String()
}

func (s *GenerateVehicleRepairPlanResponse) SetHeaders(v map[string]*string) *GenerateVehicleRepairPlanResponse {
	s.Headers = v
	return s
}

func (s *GenerateVehicleRepairPlanResponse) SetBody(v *GenerateVehicleRepairPlanResponseBody) *GenerateVehicleRepairPlanResponse {
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
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
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
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DetectObjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DetectObjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectObjectResponseBody) GoString() string {
	return s.String()
}

func (s *DetectObjectResponseBody) SetRequestId(v string) *DetectObjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectObjectResponseBody) SetData(v *DetectObjectResponseBodyData) *DetectObjectResponseBody {
	s.Data = v
	return s
}

type DetectObjectResponseBodyData struct {
	Elements []*DetectObjectResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	Width    *int32                                  `json:"Width,omitempty" xml:"Width,omitempty"`
	Height   *int32                                  `json:"Height,omitempty" xml:"Height,omitempty"`
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

func (s *DetectObjectResponseBodyData) SetWidth(v int32) *DetectObjectResponseBodyData {
	s.Width = &v
	return s
}

func (s *DetectObjectResponseBodyData) SetHeight(v int32) *DetectObjectResponseBodyData {
	s.Height = &v
	return s
}

type DetectObjectResponseBodyDataElements struct {
	Type  *string  `json:"Type,omitempty" xml:"Type,omitempty"`
	Boxes []*int32 `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s DetectObjectResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s DetectObjectResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *DetectObjectResponseBodyDataElements) SetType(v string) *DetectObjectResponseBodyDataElements {
	s.Type = &v
	return s
}

func (s *DetectObjectResponseBodyDataElements) SetBoxes(v []*int32) *DetectObjectResponseBodyDataElements {
	s.Boxes = v
	return s
}

func (s *DetectObjectResponseBodyDataElements) SetScore(v float32) *DetectObjectResponseBodyDataElements {
	s.Score = &v
	return s
}

type DetectObjectResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetectObjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DetectObjectResponse) SetBody(v *DetectObjectResponseBody) *DetectObjectResponse {
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

func (client *Client) DetectTransparentImageWithOptions(request *DetectTransparentImageRequest, runtime *util.RuntimeOptions) (_result *DetectTransparentImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetectTransparentImageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetectTransparentImage"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
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
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	detectTransparentImageReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	detectTransparentImageResp, _err := client.DetectTransparentImageWithOptions(detectTransparentImageReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectTransparentImageResp
	return _result, _err
}

func (client *Client) DetectVehicleICongestionWithOptions(tmpReq *DetectVehicleICongestionRequest, runtime *util.RuntimeOptions) (_result *DetectVehicleICongestionResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DetectVehicleICongestionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RoadRegions)) {
		request.RoadRegionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RoadRegions, tea.String("RoadRegions"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.PreRegionIntersectFeatures)) {
		request.PreRegionIntersectFeaturesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PreRegionIntersectFeatures, tea.String("PreRegionIntersectFeatures"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetectVehicleICongestionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetectVehicleICongestion"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ClassifyVehicleInsuranceWithOptions(request *ClassifyVehicleInsuranceRequest, runtime *util.RuntimeOptions) (_result *ClassifyVehicleInsuranceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ClassifyVehicleInsuranceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ClassifyVehicleInsurance"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
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
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	classifyVehicleInsuranceReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetectIPCObjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetectIPCObject"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetVehicleRepairPlanWithOptions(request *GetVehicleRepairPlanRequest, runtime *util.RuntimeOptions) (_result *GetVehicleRepairPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetVehicleRepairPlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetVehicleRepairPlan"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetVehicleRepairPlan(request *GetVehicleRepairPlanRequest) (_result *GetVehicleRepairPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVehicleRepairPlanResponse{}
	_body, _err := client.GetVehicleRepairPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectWhiteBaseImageWithOptions(request *DetectWhiteBaseImageRequest, runtime *util.RuntimeOptions) (_result *DetectWhiteBaseImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetectWhiteBaseImageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetectWhiteBaseImage"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
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
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	detectWhiteBaseImageReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	detectWhiteBaseImageResp, _err := client.DetectWhiteBaseImageWithOptions(detectWhiteBaseImageReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectWhiteBaseImageResp
	return _result, _err
}

func (client *Client) DetectMainBodyWithOptions(request *DetectMainBodyRequest, runtime *util.RuntimeOptions) (_result *DetectMainBodyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetectMainBodyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetectMainBody"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
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
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	detectMainBodyReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	detectMainBodyResp, _err := client.DetectMainBodyWithOptions(detectMainBodyReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectMainBodyResp
	return _result, _err
}

func (client *Client) DetectVehicleWithOptions(request *DetectVehicleRequest, runtime *util.RuntimeOptions) (_result *DetectVehicleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetectVehicleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetectVehicle"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
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
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	detectVehicleReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	detectVehicleResp, _err := client.DetectVehicleWithOptions(detectVehicleReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectVehicleResp
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

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetectVehicleIllegalParkingResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetectVehicleIllegalParking"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RecognizeVehicleDamageWithOptions(request *RecognizeVehicleDamageRequest, runtime *util.RuntimeOptions) (_result *RecognizeVehicleDamageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RecognizeVehicleDamageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeVehicleDamage"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
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
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	recognizeVehicleDamageReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RecognizeVehicleDashboardResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeVehicleDashboard"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
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
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	recognizeVehicleDashboardReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RecognizeVehiclePartsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeVehicleParts"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
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
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	recognizeVehiclePartsReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	recognizeVehiclePartsResp, _err := client.RecognizeVehiclePartsWithOptions(recognizeVehiclePartsReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeVehiclePartsResp
	return _result, _err
}

func (client *Client) GenerateVehicleRepairPlanWithOptions(request *GenerateVehicleRepairPlanRequest, runtime *util.RuntimeOptions) (_result *GenerateVehicleRepairPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GenerateVehicleRepairPlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GenerateVehicleRepairPlan"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateVehicleRepairPlan(request *GenerateVehicleRepairPlanRequest) (_result *GenerateVehicleRepairPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateVehicleRepairPlanResponse{}
	_body, _err := client.GenerateVehicleRepairPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectObjectWithOptions(request *DetectObjectRequest, runtime *util.RuntimeOptions) (_result *DetectObjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetectObjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetectObject"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
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
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	detectObjectReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	detectObjectResp, _err := client.DetectObjectWithOptions(detectObjectReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectObjectResp
	return _result, _err
}
